package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var (
	minEntryHeight     float32 = 37.000000
	payloadEntryHeight float32 = minEntryHeight * 2
	blockViewHeight    float32 = minEntryHeight*2 + payloadEntryHeight
	entryWidth         float32 = 600
)

type ChainLayout struct {
}

func (c *ChainLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize((entryWidth+float32(100))*float32(len(objects)), blockViewHeight)
}

func (c *ChainLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)
	for _, o := range objects {
		size := fyne.NewSize(entryWidth+float32(50), blockViewHeight)
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(size.Width+theme.Padding()*10, 0))
	}
}
