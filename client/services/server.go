package services

import (
	"fyne.io/fyne/widget"
	"github.com/gangjun06/gChat/client/lib/db"
)

type ServerInfo struct {
	Layout *widget.Box
	Info   []*db.ServerInfo
}

var ServerView *ServerInfo

func SetServerView(Layout *widget.Box) {
	ServerView = &ServerInfo{Layout: Layout}
}

func (s *ServerInfo) Refresh() {
}
