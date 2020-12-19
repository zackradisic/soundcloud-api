package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func main() {
	start := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(start)
	sc, err := soundcloudapi.New("", &http.Client{
		Timeout: time.Second * 20,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(sc.ClientID())

	playlist, err := sc.GetPlaylistInfo("https://soundcloud.com/sdlfjsdfl")

	if err != nil {
		log.Fatal(err.Error())
	}

	end := time.Now().UnixNano() / int64(time.Millisecond)
	elapsed := float32(end-start) / 1000.0
	fmt.Printf("Elapsed: %f\n", elapsed)

	fmt.Println("Playlist title: " + playlist.Title)
	fmt.Printf("Playlist length: %d\n", len(playlist.Tracks))

	for i, track := range playlist.Tracks {
		fmt.Printf("%d. %s : %d\n", i+1, track.Title, track.ID)
	}
}
