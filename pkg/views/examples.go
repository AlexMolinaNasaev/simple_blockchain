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
		"whoami":     {"Кто я", "whoami", Whoami},
		"definition": {"Определение", "definition", Definition},
		"hash":       {"Хэш", "hash", Hash},
		"block":      {"Блок", "block", Block},
		"chain":      {"Цепь", "chain", Chain},
		"mining":     {"Майнинг", "mining", Mining},
		"links":      {"Полезные ссылки", "links", Links},
	}

	ExampleIndex = map[string][]string{
		"": {"whoami", "definition", "hash", "block", "chain", "mining", "links"},
	}
)
