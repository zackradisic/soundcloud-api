package main

import (
	"fmt"
	"log"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func main() {
	sc, err := soundcloudapi.New("")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(sc.ClientID())

	query, err := sc.GetLikes(soundcloudapi.GetLikesOptions{
		ProfileURL: "https://soundcloud.com/dlfsldkjf",
		Limit:      100,
		Offset:     0,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, track := range query.Collection {
		fmt.Printf("%d. %s %s\n", i+1, track.Track.Title, track.Track.Kind)
	}
}
