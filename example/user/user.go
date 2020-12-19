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

	user, err := sc.GetUser(soundcloudapi.GetUserOptions{
		ProfileURL: "https://soundcloud.com/sisjdfg;kljs;dkg",
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(user.Username)
}
