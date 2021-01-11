package main

import (
	"fmt"
	"log"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func main() {
	sc, err := soundcloudapi.New(soundcloudapi.APIOptions{})

	if err != nil {
		log.Fatal(err.Error())
	}

	response, err := sc.Search(soundcloudapi.SearchOptions{
		Query: "redbone",
		Kind:  soundcloudapi.KindTrack,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tracks, err := response.GetTracks()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, item := range tracks {
		fmt.Println(item.Title)
	}

}
