package blockchain

import "fmt"

type Peer struct {
	ID    uint8
	Peers map[uint8]*Peer
	chain *Chain
}

func NewPeer(ID uint8) *Peer {
	return &Peer{
		ID: ID,
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
	return p.chain
}

func (p *Peer) GetChainLen() int {
	return len(p.chain.Blocks)
}

func (p *Peer) DeletePeerByID(id uint8) {
	delete(p.Peers, id)
}

func (p *Peer) AddBlock(block Block) error {
	p.chain.AddBlock(block)
	return nil
}

func (p *Peer) MineBlock(payload string) {
	chainLen := len(p.chain.Blocks)
	prevBlock := p.chain.Blocks[chainLen-1]

	newBlock := Block{
		Number:        chainLen,
		PrevBlockHash: prevBlock.Hash,
		Payload:       payload,
	}
	newBlock.Mine()

	p.chain.AddBlock(newBlock)

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

		err := largestChainPeer.chain.ValidateChain()
		if err != nil {
			p.DeletePeerByID(largestChainPeer.ID)
			if len(p.Peers) == 0 {
				return NewBlockchainChainError(PeerSyncError, nil)
			}
			continue
		}

		break
	}

	p.chain = largestChainPeer.chain
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
