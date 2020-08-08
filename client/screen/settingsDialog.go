package screen

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/gangjun06/gChat/client/services"
)

func SettingsDialog(w fyne.Window) {

	data := services.GetUserInfo()

	inputName := widget.NewEntry()
	inputName.SetPlaceHolder("John Doe..")
	inputName.SetText(data.Username)

	inputAvatar := widget.NewEntry()
	inputAvatar.SetPlaceHolder("Avatar URL....")
	inputAvatar.SetText(data.Avatar)

	form := widget.NewForm(
		widget.NewFormItem("Name", inputName),
		widget.NewFormItem("Avatar", inputAvatar),
	)

	dialog.ShowCustomConfirm("Settings", "Update", "Cancel", form, func(b bool) {
		if !b {
			return
		}

		services.SetUserInfo(inputName.Text, inputAvatar.Text)
	}, w)
}
