package views

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Links(_ fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Bitcoin block explorers"),
		makeLink("https://explorer.bitcoin.com/btc", ""),
		makeLink("https://www.blockchain.com/explorer", ""),
		makeLink("https://blockstream.info/", ""),
		widget.NewSeparator(),
		widget.NewLabel("Ethereum block explorers"),
		makeLink("https://www.blockchain.com/explorer?view=eth", ""),
		makeLink("https://ethblockexplorer.org/", ""),
		makeLink("https://www.etherchain.org/", ""),
	)
}

func makeLink(uri, description string) fyne.CanvasObject {
	link, err := url.Parse(uri)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}
	hyperlink := widget.NewHyperlink(uri, link)

	return container.NewHBox(hyperlink, widget.NewLabel(description))
}
