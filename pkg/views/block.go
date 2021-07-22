package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/alexmolinanasaev/simple_blockchain/pkg/blockchain"
)

func Block(_ fyne.Window) fyne.CanvasObject {
	peer := blockchain.NewPeer(1, 1)
	peer.MineBlock("test!")
	controller := chainViewController{
		peer: peer,
	}

	peerContent := controller.makeBlock(peer.GetBlock(1), false, false)
	return container.NewCenter(peerContent)
}
