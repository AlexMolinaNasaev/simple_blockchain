package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var (
	minEntryHeight     float32 = 37.000000
	payloadEntryHeight float32 = minEntryHeight * 2
	blockViewHeight    float32 = minEntryHeight*3 + payloadEntryHeight
	entryWidth         float32 = 600
)

type BlockLayout struct {
}

func (c *BlockLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize((entryWidth+float32(200))*float32(len(objects)), blockViewHeight)
}

func (c *BlockLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)
	for _, o := range objects {
		size := fyne.NewSize(entryWidth+float32(150), blockViewHeight)
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(size.Width+theme.Padding()*10, 0))
	}
}
