package views

import (
	"crypto/sha256"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Mining(w fyne.Window) fyne.CanvasObject {
	nonce := widget.NewEntry()
	payload := widget.NewMultiLineEntry()
	hash := widget.NewLabel("")

	calcHash := func(s string) {
		h := sha256.New()
		data := fmt.Sprintf("%s%s", nonce.Text, payload.Text)
		h.Write([]byte(data))
		hash.SetText(fmt.Sprintf("%x", h.Sum(nil)))
	}

	nonce.OnChanged = calcHash
	payload.OnChanged = calcHash
	payload.Text = "mine me!"
	nonce.SetPlaceHolder("29748")

	info := container.NewVBox(
		widget.NewLabel("nonce"),
		widget.NewLabel("payload"),
		widget.NewLabel(""),
		widget.NewLabel("hash"),
	)

	blockContent := container.NewVBox(nonce, payload, hash)

	blockContent = container.New(&BlockLayout{}, blockContent)

	content := container.NewBorder(widget.NewLabel("hash puzzle 0000"), nil, info, blockContent)

	return container.NewCenter(content)
}
