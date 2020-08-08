package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"github.com/gangjun06/gChat/client/lib/db"
	"github.com/gangjun06/gChat/client/screen"
)

func main() {
	go db.InitDB()

	a := app.NewWithID("dev.gangjun.gchat.client")
	w := a.NewWindow("gChat")
	a.Settings().SetTheme(theme.LightTheme())

	w.SetContent(screen.HomeScreen(w))
	w.Resize(fyne.NewSize(360, 560))
	w.ShowAndRun()
}
