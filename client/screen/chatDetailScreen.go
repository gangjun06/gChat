package screen

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func ChatDetailScreen() fyne.CanvasObject {
	entryLog := widget.NewLabel("")
	entryLog.Wrapping = fyne.TextWrapWord
	entryLogScroller := widget.NewVScrollContainer(entryLog)
	entry := widget.NewEntry()

	bottom := widget.NewHBox(
		entry,
		widget.NewButton("Send", func() {
			fmt.Println(entry.Text)
		}),
	)

	borderLayout := layout.NewBorderLayout(nil, nil, nil, bottom)
	return fyne.NewContainerWithLayout(
		borderLayout, bottom, entryLogScroller,
	)
}
