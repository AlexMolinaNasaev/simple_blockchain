package views

import (
	"crypto/sha256"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Hash(w fyne.Window) fyne.CanvasObject {
	input := widget.NewEntry()
	input.SetPlaceHolder("Введите текст...")
	hashResult := widget.NewEntry()
	hashResult.Disable()
	input.OnChanged = func(s string) {
		h := sha256.New()
		h.Write([]byte(input.Text))
		hashResult.SetText(fmt.Sprintf("%x", h.Sum(nil)))
	}

	field1 := widget.NewEntry()
	field2 := widget.NewEntry()

	field1.SetPlaceHolder("Введите текст...")
	field2.SetPlaceHolder("Введите текст...")

	compareHashesButton := widget.NewButton("сравнить хэши", func() {
		if field1.Text == field2.Text {
			dialog.ShowInformation("Успех!", "Хэши совпадают", w)
		} else {
			dialog.ShowInformation("Ошибка!", "Хэши не совпадают", w)
		}
	})

	hashesQ := widget.NewLabel("79228162514264337593543950336")

	content := container.NewVBox(
		input, hashResult, field1, field2,
		compareHashesButton, widget.NewLabel(""), hashesQ)

	return container.NewCenter(container.New(&BlockLayout{}, content))
}
