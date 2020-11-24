package main

import (
	"fmt"
	"log"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func main() {
	sc, err := soundcloudapi.NewSoundCloudAPI("")

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

	for _, track := range query.Collection {
		fmt.Println(track.Track.Title)
	}
}
