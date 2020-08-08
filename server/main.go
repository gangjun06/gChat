package main

import (
	"log"
	"net/http"
)

func main() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/ws", server)
	log.Fatalln(http.ListenAndServe(":5000", nil))
}
