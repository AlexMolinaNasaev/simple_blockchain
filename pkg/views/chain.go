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

	peerContent := makePeer(peer)

	scroller := container.NewHScroll(container.NewCenter(peerContent))

	return scroller
}

func makePeer(peer *blockchain.Peer) fyne.CanvasObject {
	peerChain := makeChain(peer.GetChain())
	peerContent := container.New(&BlockLayout{}, peerChain...)
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
		widget.NewTextGrid(),
	)

	return container.NewHBox(InfoContent, widget.NewSeparator(), blockContent)
}

// !TODO добавить валидацию цепи и окрашивание полей в красный от сломанного блока. Надо использовать горутину, которая будет ловить изменение пэйлоуда
// через глобалный канал
// !TODO добавить майнинг блоков
