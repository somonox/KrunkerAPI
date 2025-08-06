package KrunkerAPI

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack"
)

type Profile struct {
	Name               string
	Clan               string
	Kills              uint16
	Deaths             uint16
	Score              uint32
	Time               uint32
	Played             uint16
	Wins               uint16
	Losses             uint16
	Nukes              float64
	KR                 uint8
	Inventory          uint16
	Junk               string
	Shots              float64
	Hits               float64
	Misses             float64
	WallBangs          float64
	DateNew            string
	Followed           int8
	Following          int8
	Crouches           float64
	HeadShots          float64
	LegShots           float64
	MeleeKills         float64
	ThrowingMeleeKills float64
	FistKills          float64
	Sprays             float64
}

func (api KrunkerAPI) GetProfile(username string) (*Profile, *[]interface{}) {
	p := Profile{}

	message := []interface{}{"r", "profile", username}
	packedMessage, err := msgpack.Marshal(message)
	if err != nil {
		log.Println("Error encoding message:", err)
	}
	err = api.conn.WriteMessage(websocket.BinaryMessage, append(packedMessage, 0x00, 0x00))
	if err != nil {
		log.Println("Write error:", err)
		return nil, nil
	}

	log.Println("Sent:", message)

	// 잠시 대기
	time.Sleep(2 * time.Second)

	for {
		_, message, err := api.conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return nil, nil
		}

		// 수신한 메시지 출력
		var decodedMessage []interface{}
		err = msgpack.Unmarshal(message[:len(message)-2], &decodedMessage)
		if err != nil {
			log.Println("Error decoding message:", err)
			continue
		}

		if len(decodedMessage) > 2 {
			var player_stats map[string]interface{}
			var profile map[string]interface{}

			if decodedMessage[3] == nil {
				return nil, nil
			}

			err := json.Unmarshal([]byte(decodedMessage[3].(map[string]interface{})["player_stats"].(string)), &player_stats)
			profile = decodedMessage[3].(map[string]interface{})

			if err != nil {
				log.Println("Error: ", err)
				return nil, nil
			}

			p.Name = profile["player_name"].(string)
			p.Clan = profile["player_clan"].(string)
			p.Kills = profile["player_kills"].(uint16)
			p.Deaths = profile["player_deaths"].(uint16)
			p.Score = profile["player_score"].(uint32)
			p.Time = profile["player_timeplayed"].(uint32)
			p.Played = profile["player_games_played"].(uint16)
			p.Wins = profile["player_wins"].(uint16)
			p.Losses = p.Played - p.Wins
			p.Nukes = player_stats["n"].(float64)
			p.KR = profile["player_funds"].(uint8)
			p.Inventory = profile["player_skinvalue"].(uint16)
			p.Junk = profile["player_elo4"].(string)
			p.Shots = player_stats["s"].(float64)
			p.Hits = player_stats["hs"].(float64)
			p.Misses = p.Shots - p.Hits
			p.WallBangs = player_stats["wb"].(float64)
			p.HeadShots = player_stats["h"].(float64)
			p.LegShots = player_stats["ls"].(float64)
			p.DateNew = profile["player_datenew"].(string)
			p.Followed = profile["player_followed"].(int8)
			p.Following = profile["player_following"].(int8)
			p.Crouches = player_stats["crc"].(float64)
			p.MeleeKills = player_stats["mk"].(float64)
			p.ThrowingMeleeKills = player_stats["tmk"].(float64)
			p.FistKills = player_stats["fk"].(float64)
			p.Sprays = player_stats["spry"].(float64)

			return &p, &decodedMessage
		}
	}
}
