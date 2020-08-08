package main

import (
	"fmt"
	"net/url"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/cmd/fyne_demo/data"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func WelcomeScreen(a fyne.App) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(data.FyneScene)
	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(171, 125))
	} else {
		logo.SetMinSize(fyne.NewSize(228, 167))
	}

	return widget.NewVBox(
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Welcome to the Fyne toolkit demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),

		widget.NewHBox(layout.NewSpacer(),
			widget.NewHyperlink("fyne.io", parseURL("https://fyne.io/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("documentation", parseURL("https://fyne.io/develop/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("sponsor", parseURL("https://github.com/sponsors/fyne-io")),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),

		widget.NewGroup("Theme",
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				widget.NewButton("Dark", func() {
					a.Settings().SetTheme(theme.DarkTheme())
				}),
				widget.NewButton("Light", func() {
					a.Settings().SetTheme(theme.LightTheme())
				}),
			),
		),
	)
}

func LogScreen(a fyne.App) fyne.CanvasObject {
	entryLog := widget.NewLabel("")
	entryLog.SetText("LogFile")
	entryLog.Wrapping = fyne.TextWrapWord
	entryLogScroller := widget.NewVScrollContainer(entryLog)

	top := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			fmt.Println("Save log")
		}),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			entryLog.SetText("")
		}),
	)

	borderLayout := layout.NewBorderLayout(top, nil, nil, nil)
	return fyne.NewContainerWithLayout(
		borderLayout, top, entryLogScroller,
	)
}

func main() {
	a := app.NewWithID("dev.gangjun.gchat.server")

	w := a.NewWindow("gChat - Server")

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Home", theme.HomeIcon(), WelcomeScreen(a)),
		widget.NewTabItemWithIcon("Log", theme.FileTextIcon(), LogScreen(a)),
	)

	tabs.SetTabLocation(widget.TabLocationLeading)
	w.SetContent(tabs)

	w.ShowAndRun()
}
