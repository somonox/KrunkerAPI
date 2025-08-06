package krunkerapi

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
	headers := http.Header{"Origin": []string{"https://krunker.io"}}
	log.Printf("Websocket connection opned")

	conn, _, err := websocket.DefaultDialer.Dial(url.String(), headers)

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
