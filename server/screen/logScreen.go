package screen

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/gangjun06/gChat/server/util"
)

func LogScreen(a fyne.App) fyne.CanvasObject {
	logText := "gChat(Server) LogFile"
	entryLog := widget.NewLabel("")
	util.MainLog = util.NewLog(entryLog)
	entryLog.SetText(logText)
	entryLog.Wrapping = fyne.TextWrapWord
	entryLogScroller := widget.NewVScrollContainer(entryLog)

	top := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
		}),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			entryLog.SetText(logText)
		}),
	)

	borderLayout := layout.NewBorderLayout(top, nil, nil, nil)
	return fyne.NewContainerWithLayout(
		borderLayout, top, entryLogScroller,
	)
}
