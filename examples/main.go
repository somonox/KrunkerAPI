package main

import (
	"log"

	"github.com/somonox/KrunkerAPI"
)

func main() {
	api, err := KrunkerAPI.NewKrunkerAPI()
	if err != nil {
		log.Fatal(err)
	}
	defer api.Close()

	profile, rawData, _ := api.GetProfile("sydney")
	if profile == nil {
		log.Fatal("Failed to get profile")
	}

	log.Println("Profile:", *profile)
	log.Println("Raw data:", *rawData)
}
