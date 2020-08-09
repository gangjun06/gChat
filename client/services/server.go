package services

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/gangjun06/gChat/client/lib/db"
)

type ServerInfo struct {
	Layout *widget.Box
	Info   *[]db.ServerInfo
}

var ServerView *ServerInfo

func SetServerView(layout *widget.Box) {
	ServerView = &ServerInfo{Layout: layout}
	ServerView.Refresh()
}

func (s *ServerInfo) Refresh() {
	var data []db.ServerInfo
	userinfo := GetUserInfo()
	result := db.DB().Find(&data)
	if result.Error != nil {
		log.Fatalln(result.Error)
		return
	}
	ServerView.Info = &data
	var emptyData []fyne.CanvasObject
	s.Layout.Children = emptyData

	for _, info := range *s.Info {
		id := info.ID
		address := info.Address
		name := info.Name
		infobox := widget.NewHBox(widget.NewLabel(info.Name), layout.NewSpacer(), widget.NewLabel(info.Address))
		actionbox := widget.NewHBox(layout.NewSpacer(), widget.NewButton("Delete", func() {
			s.DeleteItem(id)
		}), widget.NewButton("Open", func() {
			labelLog := widget.NewLabel("Start Chatting(" + name + ")\n")

			socket := NewSocket(address, "8080", &userinfo, labelLog)
			go socket.Connect()

			w := fyne.CurrentApp().NewWindow("Chat Detail")
			w.Resize(fyne.NewSize(360, 560))

			labelLog.Wrapping = fyne.TextWrapWord
			labelLogScroller := widget.NewVScrollContainer(labelLog)
			entry := widget.NewEntry()

			btnSend := widget.NewButton("Send", func() {
				socket.Send(entry.Text)
				entry.SetText("")
			})

			bottomLayout := layout.NewBorderLayout(nil, nil, nil, btnSend)
			bottom := fyne.NewContainerWithLayout(
				bottomLayout,
				btnSend,
				entry,
			)

			borderLayout := layout.NewBorderLayout(nil, bottom, nil, nil)
			w.SetContent(fyne.NewContainerWithLayout(
				borderLayout, bottom, labelLogScroller))
			w.Show()

		}))
		box := widget.NewVBox(infobox, actionbox)
		s.Layout.Append(box)
	}
	s.Layout.Refresh()
}

func (s *ServerInfo) AddItem(name string, address string) {
	db.DB().Create(&db.ServerInfo{Name: name, Address: address})
	s.Refresh()
}

func (s *ServerInfo) DeleteItem(id int) {
	db.DB().Where("id = ?", id).Delete(&db.ServerInfo{})
	s.Refresh()
}
