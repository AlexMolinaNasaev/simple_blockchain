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
		"hash":   {"Хэш", "hash", Hash},
		"block":  {"Блок", "block", Block},
		"chain":  {"Цепь", "chain", Chain},
		"mining": {"Майнинг", "mining", Mining},
		"links":  {"Полезные ссылки", "links", Links},
	}

	ExampleIndex = map[string][]string{
		"": {"hash", "block", "chain", "mining", "links"},
	}
)
