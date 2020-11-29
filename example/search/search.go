package main

import (
	"encoding/json"
	"fmt"
	"log"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func main() {
	sc, err := soundcloudapi.New("")

	if err != nil {
		log.Fatal(err.Error())
	}

	response, err := sc.Search(soundcloudapi.SearchOptions{
		Query: "redbone",
		Kind:  soundcloudapi.SearchKindTrack,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, err := json.Marshal(response)

	trackQuery := &soundcloudapi.PaginatedTrackQuery{}

	err = json.Unmarshal(data, trackQuery)

	for _, item := range trackQuery.Collection {
		fmt.Println(item.Title)
	}
}
