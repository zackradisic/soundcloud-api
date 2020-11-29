package soundcloudapi_test

import (
	"encoding/json"
	"testing"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func TestSearch(t *testing.T) {

	response, err := api.Search(soundcloudapi.SearchOptions{
		Query: "redbone",
		Kind:  soundcloudapi.SearchKindTrack,
	})

	if err != nil {
		t.Error(err.Error())
		return
	}

	data, err := json.Marshal(response)
	trackQuery := &soundcloudapi.PaginatedTrackQuery{}

	err = json.Unmarshal(data, trackQuery)

	for _, item := range trackQuery.Collection {
		if item.Kind != "track" {
			t.Errorf("Not every item in PaginatedTrackQuery has kind = 'track'. Received (%s) Expected (%s) ", item.Kind, "track")
		}
	}
}
