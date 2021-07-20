package views

import (
	"crypto/sha256"
	"fmt"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Mining(w fyne.Window) fyne.CanvasObject {
	content := container.NewVBox(nonce(), widget.NewSeparator(), mining())
	return container.NewCenter(content)
}

func nonce() fyne.CanvasObject {
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

	return container.NewBorder(widget.NewLabel("hash puzzle 0000"), nil, info, blockContent)
}

func mining() fyne.CanvasObject {
	puzzleDifficulty := widget.NewEntry()
	puzzleDifficulty.SetText("0")
	nonceLabel := widget.NewLabel("nonce")
	payloadEntry := widget.NewEntry()
	payloadEntry.SetText("MINE MEEEEEEEE!")
	hashLabel := widget.NewLabel("hash")

	mineButton := widget.NewButton("mine", func() {
		re, _ := regexp.Compile(fmt.Sprintf(`^0{%s}`, puzzleDifficulty.Text))

		var nonce int

		payload := payloadEntry.Text

		for {
			h := sha256.New()
			data := fmt.Sprintf("%d%s", nonce, payload)
			h.Write([]byte(data))
			hashHex := fmt.Sprintf("%x", h.Sum(nil))

			nonceLabel.SetText(fmt.Sprintf("%d", nonce))
			hashLabel.SetText(hashHex)

			if re.Match([]byte(hashHex)) {
				return
			}
			nonce += 1
		}
	})

	content := container.NewVBox(
		puzzleDifficulty,
		nonceLabel,
		payloadEntry,
		hashLabel,
		mineButton,
	)

	return content
}
