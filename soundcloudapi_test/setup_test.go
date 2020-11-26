package soundcloudapi_test

import (
	"log"
	"os"
	"testing"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

var api *soundcloudapi.API

func TestMain(m *testing.M) {
	var err error
	api, err = soundcloudapi.New("")
	if err != nil {
		log.Fatalf("failed to create new API: %+v\n", err)
	}

	code := m.Run()
	os.Exit(code)
}
