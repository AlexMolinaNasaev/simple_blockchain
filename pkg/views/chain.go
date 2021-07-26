package views

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alexmolinanasaev/simple_blockchain/pkg/blockchain"
)

type chainViewController struct {
	peer    *blockchain.Peer
	content *fyne.Container
	hashes  []*widget.TextGrid
}

func Chain(_ fyne.Window) fyne.CanvasObject {
	peer := blockchain.NewPeer(1, 1)
	controller := chainViewController{
		peer:   peer,
		hashes: make([]*widget.TextGrid, 0),
	}
	peer.MineBlock(fmt.Sprintf("%s\n%s",
		"Nikolay -> Peter 500",
		"Nikolay -> Natalya 200"))
	peer.MineBlock(fmt.Sprintf("%s\n%s",
		"Nikolay -> Peter 500",
		"Nikolay -> Natalya 200"))
	peer.MineBlock(fmt.Sprintf("%s\n%s",
		"Olga -> Peter 111",
		"Peter -> Nikolay 14"))
	peer.MineBlock(fmt.Sprintf("%s\n%s",
		"Alex -> Vladislav 100",
		"Dominic -> Toretto 777"))
	peer.MineBlock(fmt.Sprintf("%s\n%s",
		"Nikolay -> Peter 500",
		"Nikolay -> Natalya 200"))
	peer.MineBlock(fmt.Sprintf("%s\n%s",
		"Nikolay -> Peter 500",
		"Nikolay -> Natalya 200"))

	controller.content = container.New(&BlockLayout{}, controller.makeChain()...)
	content := container.NewVBox(
		widget.NewSeparator(),
		controller.content,
		widget.NewLabel(""),
		widget.NewLabel(""),
		widget.NewSeparator(),
	)

	return container.NewHScroll(container.NewCenter(content))
}

func (c *chainViewController) makeChain() []fyne.CanvasObject {
	var blocks []fyne.CanvasObject
	for _, block := range c.peer.GetChain().Blocks {
		blocks = append(blocks, c.makeBlock(block, false, false))
	}

	emptyBlock := blockchain.Block{
		Number:        c.peer.GetChainLen(),
		PrevBlockHash: c.peer.GetBlock(c.peer.GetChainLen() - 1).Hash,
	}
	blocks = append(blocks, c.makeBlock(emptyBlock, true, false))

	return blocks
}

func (c *chainViewController) makeBlock(block blockchain.Block, minable, isWrong bool) fyne.CanvasObject {
	blockNumber := widget.NewLabel(fmt.Sprintf("%d", block.Number))

	prevBlockHash := widget.NewTextGridFromString(block.PrevBlockHash)
	prevBlockHash.SetStyleRange(0, 0, 0, 64,
		&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}})

	c.hashes = append(c.hashes, prevBlockHash)

	payload := widget.NewMultiLineEntry()
	payload.SetPlaceHolder("Введите текст")
	payload.SetText(block.Payload)

	currBlockHash := widget.NewLabel(block.Hash)
	currBlockHash.SetText(block.Hash)

	payload.OnChanged = func(s string) {
		block.Payload = payload.Text
		currBlockHash.SetText(block.Mine())

		if block.Number < c.peer.GetChainLen() {
			c.peer.Chain.Blocks[block.Number] = block
			blockNum, err := c.peer.GetChain().ValidateChain()
			if err != nil {
				for i := range c.hashes {
					if i > c.peer.GetChainLen() {
						break
					}

					if i >= blockNum {
						c.hashes[i].SetStyleRange(0, 0, 0, 64,
							&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 192, G: 64, B: 64, A: 128}})
					} else {
						c.hashes[i].SetStyleRange(0, 0, 0, 64,
							&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}})
					}

					c.hashes[i].Refresh()
				}
			}
		}
	}

	mineButton := widget.NewButton("Mine", func() {})

	mineButton.OnTapped = func() {
		c.peer.MineBlock(payload.Text)

		emptyBlock := blockchain.Block{
			Number:        c.peer.GetChainLen(),
			PrevBlockHash: c.peer.GetBlock(c.peer.GetChainLen() - 1).Hash,
		}

		emptyBlock.Mine()

		c.content.Objects = append(c.content.Objects, c.makeBlock(emptyBlock, true, false))
		mineButton.DisableableWidget.Disable()
		c.content.Refresh()
	}

	if !minable {
		mineButton.DisableableWidget.Disable()
	}

	blockContent := container.NewVBox(blockNumber, prevBlockHash, payload, currBlockHash)
	InfoContent := container.NewVBox(
		widget.NewLabel("Block number"),
		widget.NewLabel("Prev hash"),
		widget.NewLabel("Payload"),
		widget.NewLabel(""),
		widget.NewLabel("Hash"),
		mineButton,
	)

	return container.NewHBox(InfoContent, widget.NewSeparator(), blockContent)
}
