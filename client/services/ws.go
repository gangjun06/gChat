package services

import (
	"log"
	"net/url"
	"time"

	"fyne.io/fyne/widget"
	"github.com/gangjun06/gChat/client/lib/db"
	"github.com/gorilla/websocket"
)

type Socket struct {
	Address   string
	Port      string
	connected bool
	interrupt chan struct{}
	UserInfo  *db.UserInfo
	Message   chan string
	Text      *widget.Label
}

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func NewSocket(address, port string, userInfo *db.UserInfo, text *widget.Label) *Socket {
	return &Socket{Address: address, Port: port, connected: false, UserInfo: userInfo, Text: text}
}

func (s *Socket) Connect() {
	if s.connected {
		return
	}
	s.connected = true
	s.interrupt = make(chan struct{}, 1)
	s.Message = make(chan string)

	u := url.URL{Scheme: "ws", Host: s.Address + ":" + s.Port, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			var msg Message
			err := c.ReadJSON(&msg)
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", msg.Message)
			s.Text.SetText(s.Text.Text + "\n" + msg.Username + ": " + msg.Message)
		}
	}()

	for {
		select {
		case <-done:
			return
		case msg := <-s.Message:
			err := c.WriteJSON(Message{Email: "asdf", Username: s.UserInfo.Username, Message: msg})
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-s.interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func (s *Socket) Send(chat string) {
	if s.connected {
		s.Message <- chat
	}
}

func (s *Socket) Disconnect() {
	if s.connected {
		s.connected = false
		close(s.interrupt)
	}
}
