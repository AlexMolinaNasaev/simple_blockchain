package blockchain

import "fmt"

type Peer struct {
	ID    uint8
	Peers map[uint8]*Peer
	Chain *Chain
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

func (p *Peer) AddBlock(payload string) {
	// p.Chain.AddBlock(payload)
}

func (p *Peer) Broadcast() {}

func (p *Peer) Sync() {
	// largestChain :=
	// for _, peer := range p.Peers {

	// }
}
