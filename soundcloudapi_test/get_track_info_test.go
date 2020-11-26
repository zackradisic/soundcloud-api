package soundcloudapi_test

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func TestGetTrackInfoURL(t *testing.T) {
	url := "https://soundcloud.com/taliya-jenkins/double-cheese-burger-hold-the"

	tracks, err := getTrackInfo(soundcloudapi.GetTrackInfoOptions{URL: url})
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(tracks) != 1 {
		t.Error("GetTrackInfo() returned more than one track")
		return
	}
}

func getTrackInfo(options soundcloudapi.GetTrackInfoOptions) ([]soundcloudapi.Track, error) {
	tracks, err := api.GetTrackInfo(options)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to retrieve track information")
	}

	if tracks[0].Kind != "track" {
		return nil, errors.Wrap(err, fmt.Sprintf("GetTrackInfo() returned wrong type of resource: received (%s) expected (track)", tracks[0].Kind))
	}

	return tracks, nil
}

func TestGetTrackInfoIDs(t *testing.T) {
	ids := []int64{122144511, 929590315}
	tracks, err := getTrackInfo(soundcloudapi.GetTrackInfoOptions{ID: ids})
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(tracks) != len(ids) {
		t.Errorf("Received wrong amount of tracks: received (%d) expected (%d)", len(tracks), len(ids))
		return
	}
}
