package blockchain

import "fmt"

type Peer struct {
	ID    int8
	Peers map[int8]*Peer
	Chain *Chain
}

func NewPeer(ID int8) *Peer {
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

func (p *Peer) Sync() {
	// largest
	// for _, peer := range p.Peers {

	// }
}
