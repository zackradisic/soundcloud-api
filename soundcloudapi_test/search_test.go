package soundcloudapi_test

import (
	"testing"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func TestSearch(t *testing.T) {

	response, err := api.Search(soundcloudapi.SearchOptions{
		Query: "redbone",
		Kind:  soundcloudapi.KindTrack,
	})

	if err != nil {
		t.Error(err.Error())
		return
	}

	tracks, err := response.GetTracks()

	if len(tracks) == 0 {
		t.Error("Received no tracks")
		return
	}

	for _, track := range tracks {
		if track.Kind != "track" {
			t.Errorf("Kind mismatch expected (%s) received (%s)\n", "track", track.Kind)
		}
	}
}
