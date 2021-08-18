package moo

import (
	"github.com/gorilla/websocket"
	"github.com/han-joker/moo-layout/moo/conf"
	"log"
	"net/http"
)

func (s *server) StartWebSocket() {
	log.Println(conf.Instance().Int("websocket.in"))
	log.Println(conf.Instance().Float64("websocket.fl"))
	log.Println(conf.Instance().Float32("websocket.fl"))
	log.Println(conf.Instance().IntSlice("websocket.ia"))
	log.Println(conf.Instance().Float32Slice("websocket.fa"))
	log.Println(conf.Instance().StringSlice("websocket.sa"))
	log.Println(conf.Instance().BoolMap("websocket.bm"))
	log.Println(conf.Instance().IntMap("websocket.im"))
	log.Println(conf.Instance().Float64Map("websocket.fm"))
	log.Println(conf.Instance().String("websocket.addr"))
	//http.HandleFunc("/", echo)
	//log.Fatal(http.ListenAndServe(conf.Instance().String("websocket.addr"), nil))
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
