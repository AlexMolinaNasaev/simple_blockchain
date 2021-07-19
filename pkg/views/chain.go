package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alexmolinanasaev/simple_blockchain/pkg/blockchain"
)

func Chain(_ fyne.Window) fyne.CanvasObject {
	peer := blockchain.NewPeer(1, 1)
	peer.MineBlock("Hello, World!")
	peer.MineBlock("test")
	peer.MineBlock("This is a new block")
	peer.MineBlock("foo bar baz")

	// peer2 := blockchain.NewPeer(2, 1)

	// peer1.AddPeer(peer2)
	// peer2.AddPeer(peer1)
	// peer2.Sync()

	peerContent := makePeer(peer)
	// peer2Content := makePeer(peer2)

	// content := container.NewVBox(
	// 	peer1Content,
	// 	widget.NewLabel(""),
	// 	widget.NewLabel(""),
	// 	peer2Content,
	// )

	scroller := container.NewHScroll(container.NewCenter(peerContent))

	return scroller
}

func makePeer(peer *blockchain.Peer) fyne.CanvasObject {
	peerChain := makeChain(peer.GetChain())
	peerContent := container.New(&ChainLayout{}, peerChain...)
	peerContent = container.NewVBox(
		widget.NewSeparator(),
		peerContent,
		widget.NewLabel(""),
		widget.NewSeparator(),
	)

	return peerContent
}

func makeChain(chain *blockchain.Chain) []fyne.CanvasObject {
	var items []fyne.CanvasObject
	for _, block := range chain.Blocks {
		items = append(items, makeBlock(block))
	}
	return items
}

func makeBlock(block blockchain.Block) fyne.CanvasObject {
	blockNumber := widget.NewLabel(fmt.Sprintf("%d", block.Number))

	prevBlockHash := widget.NewLabel(block.PrevBlockHash)

	payload := widget.NewMultiLineEntry()
	payload.SetPlaceHolder("Введите текст")
	payload.SetText(block.Payload)

	currBlockHash := widget.NewLabel(block.Hash)
	currBlockHash.SetText(block.Hash)

	payload.OnChanged = func(s string) {
		block.Payload = payload.Text
		currBlockHash.SetText(block.Mine())
	}

	blockContent := container.NewVBox(blockNumber, prevBlockHash, payload, currBlockHash)
	InfoContent := container.NewVBox(
		widget.NewLabel("Block number"),
		widget.NewLabel("Prev hash"),
		widget.NewLabel("Payload"),
		widget.NewLabel(""),
		widget.NewLabel("Hash"),
	)

	return container.NewHBox(InfoContent, widget.NewSeparator(), blockContent)
}
