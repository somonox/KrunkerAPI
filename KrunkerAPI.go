package KrunkerAPI

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type KrunkerAPI struct {
	conn *websocket.Conn
}

func NewKrunkerAPI() (*KrunkerAPI, error) {
	url := url.URL{Scheme: "wss", Host: "social.krunker.io", Path: "/ws"}
	header := http.Header{}
	header.Add("Origin", "https://krunker.io")
	header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36")
	header.Add("Pragma", "no-cache")
	header.Add("Cache-Control", "no-cache")

	log.Printf("Websocket connection opened")

	conn, _, err := websocket.DefaultDialer.Dial(url.String(), header)

	if err != nil {
		log.Fatal("dial:", err)
		return nil, err
	}

	return &KrunkerAPI{conn: conn}, nil
}

func (api *KrunkerAPI) Close() {
	if api.conn != nil {
		api.conn.Close()
		log.Println("Websocket connection closed")
	}
}
