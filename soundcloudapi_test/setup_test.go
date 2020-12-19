package soundcloudapi_test

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

var api *soundcloudapi.API

func TestMain(m *testing.M) {
	var err error
	api, err = soundcloudapi.New("", &http.Client{
		Timeout: time.Second * 20,
	})
	if err != nil {
		log.Fatalf("failed to create new API: %+v\n", err)
	}

	code := m.Run()
	os.Exit(code)
}
