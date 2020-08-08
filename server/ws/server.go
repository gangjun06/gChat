package ws

import (
	"net/http"

	"github.com/gangjun06/gChat/server/util"
)

func Serve(port string) {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	util.MainLog.AddLog("Server is listening on port "+port, util.LOG_INFO)
	http.ListenAndServe(":"+port, nil)
}
