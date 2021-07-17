package views

import (
	"fyne.io/fyne/v2"
)

type Example struct {
	Title, Intro string
	View         func(window fyne.Window) fyne.CanvasObject
}

var (
	Examples = map[string]Example{
		"hash": {"Хэш", "hash", Hash},
	}

	ExampleIndex = map[string][]string{
		"": {"hash"},
	}
)
