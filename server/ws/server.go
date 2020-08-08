package ws

import (
	"context"
	"net/http"
	"time"

	"github.com/gangjun06/gChat/server/util"
	socketio "github.com/googollee/go-socket.io"
)

var Err chan error
var stop chan bool
var serverOn bool

func Serve(port string) error {
	Err = make(chan error)
	stop = make(chan bool)

	ws, err := socketio.NewServer(nil)
	server := &http.Server{Addr: ":" + port, Handler: nil}
	if err != nil {
		return err
	}

	http.Handle("/ws", ws)

	go func() {
		err := server.ListenAndServe()
		if err.Error() != "http: Server closed" {
			util.MainLog.AddLog("Server Error", util.LOG_ERROR)
			util.MainLog.AddLog(err.Error(), util.LOG_ERROR)
			stop <- false
		}
	}()

	go func() {
		util.MainLog.AddLog("Server is listening on port "+port, util.LOG_INFO)
		serverOn = true
		stop := <-stop
		serverOn = false
		if stop {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				util.MainLog.AddLog("Closing Server", util.LOG_ERROR)
				util.MainLog.AddLog(err.Error(), util.LOG_ERROR)
			}
		}
	}()
	return nil
}

func Stop() {
	if serverOn {
		stop <- true
		util.MainLog.AddLog("Server is closed", util.LOG_INFO)
	}
}
