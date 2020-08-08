package screen

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/gangjun06/gChat/client/services"
)

func HomeScreen(w fyne.Window) fyne.CanvasObject {
	top := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			AddServerDialog(w)
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			SettingsDialog(w)
		}),
	)

	services.SetServerView(widget.NewVBox())

	borderLayout := layout.NewBorderLayout(top, nil, nil, nil)
	return fyne.NewContainerWithLayout(
		borderLayout, top, services.ServerView.Layout,
	)
}
