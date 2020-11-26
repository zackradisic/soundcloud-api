package soundcloudapi_test

import (
	"fmt"
	"net/http"
	"testing"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func TestFetchClientID(t *testing.T) {
	clientID, err := soundcloudapi.FetchClientID()
	if err != nil {
		t.Errorf("Failed to fetch client ID: %+v\n", err)
		return
	}

	res, err := http.Get(fmt.Sprintf("https://api-v2.soundcloud.com/users/547647/likes?client_id=%s&limit=1", clientID))
	if err != nil {
		t.Errorf("Failed to make SoundCloud API request: %+v\n", err)
		return
	}

	if res.StatusCode == 401 {
		t.Error("Client ID is invalid")
		return
	}
}
