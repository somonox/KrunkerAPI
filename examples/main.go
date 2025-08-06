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
	
	profile, rawData := api.GetProfile("a6a6")
	if profile == nil {
		log.Fatal("Failed to get profile")
	}

	log.Println("Profile:", *profile)
	log.Println("Raw data:", *rawData)
}
