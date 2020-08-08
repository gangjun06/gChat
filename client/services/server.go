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

func SetServerView(Layout *widget.Box) {
	ServerView = &ServerInfo{Layout: Layout}
	ServerView.Refresh()
}

func (s *ServerInfo) Refresh() {
	var data []db.ServerInfo
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
		infobox := widget.NewHBox(widget.NewLabel(info.Name), layout.NewSpacer(), widget.NewLabel(info.Address))
		actionbox := widget.NewHBox(layout.NewSpacer(), widget.NewButton("Delete", func() {
			s.DeleteItem(id)
		}), widget.NewButton("Open", func() {
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
