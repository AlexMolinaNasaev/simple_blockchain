package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/alexmolinanasaev/simple_blockchain/pkg/views"
)

func main() {
	app := app.New()
	views.RunApp(app)
}
