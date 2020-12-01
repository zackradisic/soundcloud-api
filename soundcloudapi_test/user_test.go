package soundcloudapi_test

import (
	"testing"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func TestGetUser(t *testing.T) {
	user, err := api.GetUser(soundcloudapi.GetUserOptions{ProfileURL: "https://soundcloud.com/jaiseanforever"})
	if err != nil {
		t.Error(err.Error())
		return
	}

	if user.Kind != "user" {
		t.Errorf("Kind mismatch. Expected (%s) Received (%s)", "user", user.Kind)
		return
	}
}
