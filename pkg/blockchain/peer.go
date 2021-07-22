package blockchain

import "fmt"

type Peer struct {
	ID    uint8
	Peers map[uint8]*Peer
	Chain *Chain
}

func NewPeer(peerID, chainID uint8) *Peer {
	return &Peer{
		ID:    peerID,
		Chain: NewChain(chainID),
		Peers: make(map[uint8]*Peer),
	}
}

func (p *Peer) AddPeer(peer *Peer) error {
	if _, ok := p.Peers[peer.ID]; !ok {
		p.Peers[p.ID] = peer
	} else {
		return fmt.Errorf("cannot connect to peer with same id: %d", p.ID)
	}

	return nil
}

func (p *Peer) GetChain() *Chain {
	return p.Chain
}

func (p *Peer) GetBlock(blockNum int) Block {
	return p.Chain.GetBlock(blockNum)
}

func (p *Peer) GetChainLen() int {
	return len(p.Chain.Blocks)
}

func (p *Peer) DeletePeerByID(id uint8) {
	delete(p.Peers, id)
}

func (p *Peer) AddBlock(block Block) error {
	p.Chain.AddBlock(block)
	return nil
}

func (p *Peer) MineBlock(payload string) {
	chainLen := len(p.Chain.Blocks)
	prevBlock := p.GetBlock(chainLen - 1)

	newBlock := Block{
		Number:        chainLen,
		PrevBlockHash: prevBlock.Hash,
		Payload:       payload,
	}
	newBlock.Mine()

	p.Chain.AddBlock(newBlock)

	p.BroadcastBlock(newBlock)
}

func (p *Peer) BroadcastBlock(block Block) {
	for _, peer := range p.Peers {
		peer.AddBlock(block)
	}
}

func (p *Peer) Sync() error {
	var largestChainPeer *Peer
	for {
		largestChainPeer = p.getLargestChainPeer()

		_, err := largestChainPeer.Chain.ValidateChain()
		if err != nil {
			p.DeletePeerByID(largestChainPeer.ID)
			if len(p.Peers) == 0 {
				return NewBlockchainChainError(PeerSyncError, nil)
			}
			continue
		}

		break
	}

	p.Chain = largestChainPeer.Chain
	return nil
}

func (p *Peer) getLargestChainPeer() *Peer {
	var largestChainLen int
	var largestChainPeer *Peer

	for _, peer := range p.Peers {
		chainLen := peer.GetChainLen()
		if chainLen > largestChainLen {
			largestChainLen = chainLen
			largestChainPeer = peer
		}
	}

	return largestChainPeer
}
