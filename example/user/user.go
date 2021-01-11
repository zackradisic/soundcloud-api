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

	fmt.Println(sc.ClientID())

	user, err := sc.GetUser(soundcloudapi.GetUserOptions{
		ProfileURL: "https://soundcloud.com/sisjdfg;kljs;dkg",
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(user.Username)
}
