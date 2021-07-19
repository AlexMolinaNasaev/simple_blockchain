package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	minEntryHeight     float32 = 37.000000
	payloadEntryHeight float32 = minEntryHeight * 2
	blockViewHeight    float32 = minEntryHeight*2 + payloadEntryHeight
	entryWidth         float32 = 100
)

func Chain(w fyne.Window) fyne.CanvasObject {
	// chain := blockchain.NewChain(0)

	content := container.NewHBox(
		makeBlock(0), makeBlock(1), makeBlock(2), makeBlock(3), makeBlock(4), makeBlock(5),
		makeBlock(6), makeBlock(7), makeBlock(8), makeBlock(9), makeBlock(10), makeBlock(11),
		makeBlock(12), makeBlock(13), makeBlock(14), makeBlock(15), makeBlock(16), makeBlock(17),
		makeBlock(18), makeBlock(19), makeBlock(20), makeBlock(21), makeBlock(22), makeBlock(23),
		makeBlock(24), makeBlock(25), makeBlock(26), makeBlock(27), makeBlock(28), makeBlock(29),
	)
	// content.Move(fyne.NewPos(0, 250))

	scroller := container.NewHScroll(content)
	scroller.Resize(fyne.NewSize(entryWidth*float32(len(content.Objects)), blockViewHeight))
	return container.NewMax(scroller)
}

// func makeChain(chain blockchain.Chain) []fyne.CanvasObject {

// }

func makeBlock(blockNum int) fyne.CanvasObject {
	xOff := entryWidth * float32(blockNum)

	prevBlockHash := widget.NewEntry()
	prevBlockHash.Resize(fyne.NewSize(entryWidth, prevBlockHash.MinSize().Height))
	prevBlockHash.Move(fyne.NewPos(xOff, 0))
	prevBlockHash.SetText(fmt.Sprint((blockNum)))

	payload := widget.NewMultiLineEntry()
	payload.SetPlaceHolder("Введите текст")
	payload.Resize(fyne.NewSize(entryWidth, payloadEntryHeight))
	payload.Move(fyne.NewPos(xOff, prevBlockHash.MinSize().Height+theme.Padding()))

	currBlockHash := widget.NewEntry()
	currBlockHash.Disable()
	currBlockHash.Resize(fyne.NewSize(entryWidth, currBlockHash.MinSize().Height))
	currBlockHash.Move(fyne.NewPos(xOff, prevBlockHash.MinSize().Height+payload.Size().Height+theme.Padding()))

	// blockContent := container.New(layout.NewVBoxLayout(), prevBlockHash, payload, currBlockHash)
	blockContent := container.NewWithoutLayout(prevBlockHash, payload, currBlockHash)
	blockContent.Resize(fyne.NewSize(entryWidth, blockViewHeight))
	// blockContent.Move(fyne.NewPos(0, 250))
	// return blockContent
	return container.NewVBox(blockContent)
}
