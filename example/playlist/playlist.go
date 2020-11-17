package main

import (
	"fmt"
	"log"
	"time"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func main() {
	start := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(start)
	sc, err := soundcloudapi.NewSoundCloudAPI("")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(sc.ClientID())

	playlist, err := sc.GetPlaylistInfo("kdfgksdhkljhgls")

	if err != nil {
		log.Fatal(err.Error())
	}

	end := time.Now().UnixNano() / int64(time.Millisecond)
	elapsed := float32(end-start) / 1000.0
	fmt.Printf("Elapsed: %f\n", elapsed)

	fmt.Println("Playlist title: " + playlist.Title)
}
