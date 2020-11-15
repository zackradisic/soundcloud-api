package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zackradisic/soundcloud-api/pkg/soundcloudapi"
)

func main() {
	start := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(start)
	sc, err := soundcloudapi.NewSoundCloudAPI("")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(sc.ClientID())

	_, err = sc.GetPlaylistInfo("kdfgksdhkljhgls")

	if err != nil {
		log.Fatal(err.Error())
	}

	end := time.Now().UnixNano() / int64(time.Millisecond)
	elapsed := float32(end-start) / 1000.0
	fmt.Printf("Elapsed: %f\n", elapsed)
}
