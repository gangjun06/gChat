package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/gangjun06/gChat/server/screen"
)

func main() {
	a := app.NewWithID("dev.gangjun.gchat.server")
	a.Settings().SetTheme(theme.LightTheme())

	w := a.NewWindow("gChat - Server")
	w.Resize(fyne.NewSize(600, 300))

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Home", theme.HomeIcon(), screen.HomeScreen(a)),
		widget.NewTabItemWithIcon("Log", theme.FileTextIcon(), screen.LogScreen(a)),
	)

	tabs.SetTabLocation(widget.TabLocationLeading)
	w.SetContent(tabs)

	w.ShowAndRun()
}
