package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Chain(w fyne.Window) fyne.CanvasObject {
	// chain := blockchain.NewChain(0)

	hlist := makeChain(10)
	content := container.New(&BlockLayout{}, hlist...)
	scroller := container.NewHScroll(container.NewCenter(content))

	return scroller
}

func makeChain(count int) []fyne.CanvasObject {
	var items []fyne.CanvasObject
	for i := 0; i <= count; i++ {
		items = append(items, makeBlock(i))
	}
	return items
}

func makeBlock(blockNum int) fyne.CanvasObject {
	prevBlockHash := widget.NewEntry()
	prevBlockHash.SetText(fmt.Sprint((blockNum)))

	payload := widget.NewMultiLineEntry()
	payload.SetPlaceHolder("Введите текст")

	currBlockHash := widget.NewEntry()
	currBlockHash.Disable()

	return container.NewVBox(prevBlockHash, payload, currBlockHash)
}
