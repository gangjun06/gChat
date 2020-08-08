package screen

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func SettingsDialog(w fyne.Window) {

	inputName := widget.NewEntry()
	inputName.SetPlaceHolder("John Doe..")

	inputAvatar := widget.NewEntry()
	inputAvatar.SetPlaceHolder("Avatar URL....")

	form := widget.NewForm(
		widget.NewFormItem("Name", inputName),
		widget.NewFormItem("Avatar", inputAvatar),
	)

	dialog.ShowCustomConfirm("Settings", "Update", "Cancel", form, func(b bool) {
		if !b {
			return
		}

		log.Println("Please Update", inputName.Text, inputAvatar.Text)
	}, w)
}
