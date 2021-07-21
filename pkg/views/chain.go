package views

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alexmolinanasaev/simple_blockchain/pkg/blockchain"
)

type blockViewData struct {
	blockNumber   *widget.TextGrid
	prevBlockHash *widget.TextGrid
	payload       *widget.Entry
	hash          *widget.TextGrid
}

var changeChan = make(chan struct{})
var blocksData = make([]*blockViewData, 0)
var globalChain = blockchain.Chain{}

func Chain(_ fyne.Window) fyne.CanvasObject {
	peer := blockchain.NewPeer(1, 1)
	peer.MineBlock("Hello, World!")
	peer.MineBlock("test")
	peer.MineBlock("This is a new block")
	peer.MineBlock("foo bar baz")

	globalChain := peer.GetChain()

	content := container.NewVBox(
		widget.NewSeparator(),
		makeChain(globalChain),
		widget.NewLabel(""),
		widget.NewSeparator(),
	)

	go func() {
		for {
			<-changeChan
			blockNumber, err := globalChain.ValidateChain()
			if err != nil {
				for _, b := range blocksData[:blockNumber] {
					b.prevBlockHash.SetStyleRange(0, 0, 0, 63,
						&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 192, G: 64, B: 64, A: 128}})
					b.prevBlockHash.Refresh()
				}
			}
		}
	}()

	scroller := container.NewHScroll(container.NewCenter(content))

	return scroller
}

func makeChain(chain *blockchain.Chain) fyne.CanvasObject {
	var items []fyne.CanvasObject
	for _, block := range chain.Blocks {
		items = append(items, makeBlock(block))
	}
	return container.New(&BlockLayout{}, items...)
}

func makeBlock(block blockchain.Block) fyne.CanvasObject {
	blockData := &blockViewData{
		blockNumber:   widget.NewTextGridFromString(fmt.Sprintf("%d\n", block.Number)),
		prevBlockHash: widget.NewTextGridFromString(fmt.Sprintf("%s\n", block.PrevBlockHash)),
		payload:       widget.NewMultiLineEntry(),
		hash:          widget.NewTextGridFromString(fmt.Sprintf("\n%s", block.Hash)),
	}

	blockData.prevBlockHash.SetStyleRange(0, 0, 0, 63,
		&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}})

	blocksData = append(blocksData, blockData)

	blockData.payload.SetPlaceHolder("Введите текст")
	blockData.payload.SetText(block.Payload)

	blockData.payload.OnChanged = func(s string) {
		block.Payload = blockData.payload.Text
		blockData.hash.SetText(block.Mine())
		// !TODO надо изменить саму цепочку, иначе при валидации всегда будет правильно
		changeChan <- struct{}{}
	}

	blockContent := container.NewVBox(
		blockData.blockNumber,
		blockData.prevBlockHash,
		blockData.payload,
		blockData.hash,
	)
	InfoContent := container.NewVBox(
		widget.NewTextGridFromString("Block number\n"),
		widget.NewTextGridFromString("Prev hash\n"),
		widget.NewTextGridFromString("Payload\n\n\n"),
		widget.NewTextGridFromString("Hash"),
	)

	return container.NewHBox(InfoContent, widget.NewSeparator(), blockContent)
}

// !TODO добавить валидацию цепи и окрашивание полей в красный от сломанного блока. Надо использовать горутину, которая будет ловить изменение пэйлоуда
// через канал
// !TODO добавить майнинг блоков
