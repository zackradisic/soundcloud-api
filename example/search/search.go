package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func main() {
	sc, err := soundcloudapi.New("", &http.Client{
		Timeout: time.Second * 20,
	})

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
