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
	input.SetPlaceHolder("Enter text...")
	hashResult := widget.NewEntry()
	hashResult.Disable()
	input.OnChanged = func(s string) {
		h := sha256.New()
		h.Write([]byte(input.Text))
		hashResult.SetText(fmt.Sprintf("%x", h.Sum(nil)))
	}

	field1 := widget.NewEntry()
	field2 := widget.NewEntry()

	field1.SetPlaceHolder("Enter text...")
	field2.SetPlaceHolder("Enter text...")

	compareHashesButtoion := widget.NewButton("compare hashes", func() {
		if field1.Text == field2.Text {
			dialog.ShowInformation("Success", "Hashes are equal", w)
		} else {
			dialog.ShowError(fmt.Errorf("hases are not equal"), w)
		}
	})

	return container.NewVBox(input, hashResult, field1, field2, compareHashesButtoion)
}
