package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alexmolinanasaev/simple_blockchain/pkg/blockchain"
)

func Chain(w fyne.Window) fyne.CanvasObject {
	peer1 := blockchain.NewPeer(1, 1)
	peer1.MineBlock("Hello, World!")
	peer1.MineBlock("test")
	peer1.MineBlock("This is a new block")
	peer1.MineBlock("foo bar baz")

	peer2 := blockchain.NewPeer(2, 1)

	peer1.AddPeer(peer2)
	peer2.AddPeer(peer1)
	peer2.Sync()

	// peer1Chain := makeChain(peer1.GetChain())
	// peer1Content := container.New(&ChainLayout{}, peer1Chain...)
	peer1Content := makePeer(peer1)

	// peer2Chain := makeChain(peer2.GetChain())
	// peer2Content := container.New(&ChainLayout{}, peer2Chain...)
	peer2Content := makePeer(peer2)

	content := container.NewVBox(
		peer1Content,
		widget.NewLabel(""),
		widget.NewLabel(""),
		peer2Content,
	)

	scroller := container.NewHScroll(container.NewCenter(content))

	return scroller
}

func makePeer(peer *blockchain.Peer) fyne.CanvasObject {
	peerChain := makeChain(peer.GetChain())
	peerContent := container.New(&ChainLayout{}, peerChain...)
	peerContent = container.NewVBox(
		widget.NewLabel(fmt.Sprintf("Peer №%d\n", peer.ID)),
		widget.NewSeparator(),
		peerContent,
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

	blockContent := container.NewVBox(prevBlockHash, payload, currBlockHash)
	InfoContent := container.NewVBox(
		widget.NewLabel("Prev hash"),
		widget.NewLabel("Payload"),
		widget.NewLabel(""),
		widget.NewLabel("Hash"),
	)

	return container.NewHBox(InfoContent, widget.NewSeparator(), blockContent)
}
