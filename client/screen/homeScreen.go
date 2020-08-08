package screen

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func HomeScreen(w fyne.Window) fyne.CanvasObject {
	top := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			SettingsDialog(w)
		}),
	)

	borderLayout := layout.NewBorderLayout(top, nil, nil, nil)
	return fyne.NewContainerWithLayout(
		borderLayout, top, widget.NewVBox(),
	)
}
