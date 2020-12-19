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

	fmt.Println(sc.ClientID())

	query, err := sc.GetLikes(soundcloudapi.GetLikesOptions{
		ProfileURL: "https://soundcloud.com/dlfsldkjf",
		Limit:      100,
		Offset:     0,
		Type:       "track",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	likes, err := query.GetLikes()

	for _, like := range likes {
		fmt.Println(like.Track)
	}
}
