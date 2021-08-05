package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Definition(w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel(`IBM
		Блокчейн — это совместно используемый, неизменный реестр, предназначенный для записи транзакций, учета активов и построения доверительных отношений`),
		widget.NewLabel(`Wikipedia 
		Блокчейн — выстроенная по определённым правилам непрерывная последовательная цепочка блоков (связный список), содержащих информацию`),
		widget.NewLabel(`Oracle
		Блокчейн — это реестр децентрализованных данных, которыми можно безопасно обмениваться`),
		widget.NewLabel(""),
		widget.NewSeparator(),
		widget.NewLabel(""),
		widget.NewLabel("Основные особенности"),
		widget.NewLabel("- Неизменность данных"),
		widget.NewLabel("- Децентрализация"),
		widget.NewLabel("- Криптография как доказательство владения"),
		widget.NewLabel(""),
		widget.NewLabel("Дополнительные преимущества"),
		widget.NewLabel("- Прозрачность данных"),
		widget.NewLabel("- Анонимность"),
	)
}
