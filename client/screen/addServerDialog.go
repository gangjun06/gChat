package screen

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/gangjun06/gChat/client/services"
)

func AddServerDialog(w fyne.Window) {

	inputName := widget.NewEntry()
	inputName.SetPlaceHolder("Hello, world!")

	inputAddress := widget.NewEntry()
	inputAddress.SetPlaceHolder("127.0.0.1")

	form := widget.NewForm(
		widget.NewFormItem("Name", inputName),
		widget.NewFormItem("Address", inputAddress),
	)

	dialog.ShowCustomConfirm("Add Server", "Add", "Cancel", form, func(b bool) {
		if !b {
			return
		}
		services.ServerView.AddItem(inputName.Text, inputAddress.Text)
	}, w)
}
