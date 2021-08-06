package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Whoami(w fyne.Window) fyne.CanvasObject {
	content := widget.NewLabel(`
	Спикер: Молина-Насаев Александр

	Род деятельности: backend, blockchain developer

	Опыт с блокчейнами:
		- Eos.io
		- Ethereum
		- Bitcoin
		- Tron
		- Minter
		- Binance Smart Chain
		- IBM Hyperledger Fabric

	Цель:
		Ознакомить слушателей с технологией блокчейн, рассказать о преимуществах, минусах и 
		областях применения. Поделиться опытом
	`)

	return container.NewVBox(content)
}
