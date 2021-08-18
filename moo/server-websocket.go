package moo

import (
	"github.com/gorilla/websocket"
	"github.com/han-joker/moo-layout/moo/confm"
	"log"
	"net/http"
)

func (s *server) StartWebSocket() {
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(confm.Instance().String("websocket.addr"), nil))
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
