package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/alexmolinanasaev/simple_blockchain/pkg/views"
)

func main() {
	app := app.New()
	mainWindow := app.NewWindow("Entry Widget")

	mainWindow.Resize(fyne.NewSize(1000, 1000))

	content := container.NewMax()
	setExample := func(example views.Example) {
		content.Objects = []fyne.CanvasObject{example.View(mainWindow)}
		content.Refresh()
	}

	example := container.NewBorder(
		container.NewVBox(widget.NewSeparator()), nil, nil, nil, content)

	split := container.NewHSplit(makeNav(setExample, true), example)
	split.Offset = 0.2
	mainWindow.SetContent(split)
	mainWindow.ShowAndRun()
}

func makeNav(setExample func(example views.Example), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return views.ExampleIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := views.ExampleIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			e, ok := views.Examples[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(e.Title)
		},
		OnSelected: func(uid string) {
			if e, ok := views.Examples[uid]; ok {
				setExample(e)
			}
		},
	}

	themes := container.New(layout.NewGridLayout(2),
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}
