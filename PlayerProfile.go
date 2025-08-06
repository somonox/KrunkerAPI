package KrunkerAPI

import (
	"encoding/json"
	"errors"
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
	KR                 uint16
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

func toString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func toUint16(m map[string]interface{}, key string) uint16 {
	if v, ok := m[key]; ok {
		switch val := v.(type) {
		case float64:
			return uint16(val)
		case int:
			return uint16(val)
		case uint16:
			return val
		}
	}
	return 0
}

func toUint32(m map[string]interface{}, key string) uint32 {
	if v, ok := m[key]; ok {
		switch val := v.(type) {
		case float64:
			return uint32(val)
		case int:
			return uint32(val)
		case uint32:
			return val
		}
	}
	return 0
}

func toInt8(m map[string]interface{}, key string) int8 {
	if v, ok := m[key]; ok {
		switch val := v.(type) {
		case float64:
			return int8(val)
		case int:
			return int8(val)
		case int8:
			return val
		}
	}
	return 0
}

func toFloat64(m map[string]interface{}, key string) float64 {
	if v, ok := m[key]; ok {
		if f, ok := v.(float64); ok {
			return f
		}
	}
	return 0
}

func (api KrunkerAPI) GetProfile(username string) (*Profile, *[]interface{}, error) {
	message := []interface{}{"r", "profile", username}
	packedMessage, err := msgpack.Marshal(message)
	if err != nil {
		return nil, nil, errors.New("failed to encode message: " + err.Error())
	}

	err = api.conn.WriteMessage(websocket.BinaryMessage, append(packedMessage, 0x00, 0x00))
	if err != nil {
		return nil, nil, errors.New("failed to write message: " + err.Error())
	}

	log.Println("Sent:", message)

	time.Sleep(2 * time.Second)

	for {
		_, msg, err := api.conn.ReadMessage()
		if err != nil {
			return nil, nil, errors.New("failed to read message: " + err.Error())
		}

		if len(msg) < 2 {
			continue // 메시지가 너무 짧으면 무시
		}

		var decodedMessage []interface{}
		err = msgpack.Unmarshal(msg[:len(msg)-2], &decodedMessage)
		if err != nil {
			log.Println("Failed to decode message:", err)
			continue
		}

		if len(decodedMessage) <= 3 || decodedMessage[3] == nil {
			continue
		}

		profileMap, ok := decodedMessage[3].(map[string]interface{})
		if !ok {
			continue
		}

		playerStatsStr := toString(profileMap, "player_stats")
		if playerStatsStr == "" {
			continue
		}

		playerStats := map[string]interface{}{}
		if err := json.Unmarshal([]byte(playerStatsStr), &playerStats); err != nil {
			log.Println("Failed to unmarshal player_stats:", err)
			continue
		}

		p := &Profile{
			Name:               toString(profileMap, "player_name"),
			Clan:               toString(profileMap, "player_clan"),
			Kills:              toUint16(profileMap, "player_kills"),
			Deaths:             toUint16(profileMap, "player_deaths"),
			Score:              toUint32(profileMap, "player_score"),
			Time:               toUint32(profileMap, "player_timeplayed"),
			Played:             toUint16(profileMap, "player_games_played"),
			Wins:               toUint16(profileMap, "player_wins"),
			Losses:             toUint16(profileMap, "player_games_played") - toUint16(profileMap, "player_wins"),
			Nukes:              toFloat64(playerStats, "n"),
			KR:                 toUint16(profileMap, "player_funds"),
			Inventory:          toUint16(profileMap, "player_skinvalue"),
			Junk:               toString(profileMap, "player_elo4"),
			Shots:              toFloat64(playerStats, "s"),
			Hits:               toFloat64(playerStats, "hs"),
			Misses:             toFloat64(playerStats, "s") - toFloat64(playerStats, "hs"),
			WallBangs:          toFloat64(playerStats, "wb"),
			HeadShots:          toFloat64(playerStats, "h"),
			LegShots:           toFloat64(playerStats, "ls"),
			DateNew:            toString(profileMap, "player_datenew"),
			Followed:           toInt8(profileMap, "player_followed"),
			Following:          toInt8(profileMap, "player_following"),
			Crouches:           toFloat64(playerStats, "crc"),
			MeleeKills:         toFloat64(playerStats, "mk"),
			ThrowingMeleeKills: toFloat64(playerStats, "tmk"),
			FistKills:          toFloat64(playerStats, "fk"),
			Sprays:             toFloat64(playerStats, "spry"),
		}

		return p, &decodedMessage, nil
	}
}
