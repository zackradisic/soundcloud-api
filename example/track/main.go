package main

import (
	"fmt"
	"log"
	"os"
	"time"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func main() {
	start := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(start)
	sc, err := soundcloudapi.New(soundcloudapi.APIOptions{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(sc.ClientID())

	tracks, err := sc.GetTrackInfo(soundcloudapi.GetTrackInfoOptions{
		URL: "sdfkljsdflkj",
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Track title: " + tracks[0].Title)

	out, err := os.Create("output.mp3")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer out.Close()
	err = sc.DownloadTrack(tracks[0].Media.Transcodings[0], out)
	if err != nil {
		log.Fatal(err.Error())
	}

	end := time.Now().UnixNano() / int64(time.Millisecond)
	elapsed := float32(end-start) / 1000.0
	fmt.Printf("Elapsed: %f\n", elapsed)
}
