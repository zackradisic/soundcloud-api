package soundcloudapi_test

import (
	"testing"

	"github.com/pkg/errors"
	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func getPlaylistInfo(url string) (soundcloudapi.Playlist, error) {
	var playlist soundcloudapi.Playlist
	var err error
	playlist, err = api.GetPlaylistInfo(url)
	if err != nil {
		return playlist, err
	}

	if playlist.Kind != "playlist" {
		return playlist, errors.Errorf("Invalid kind property on returned playlist: received (%s) expected (%s)\n", playlist.Kind, "playlist")
	}

	return playlist, nil
}

func TestGetPlaylistInfoLessThan55Tracks(t *testing.T) {
	_, err := getPlaylistInfo("https://soundcloud.com/ilyanaazman/sets/latenightlofi")
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestGetPlaylistInfoGreaterThan55Tracks(t *testing.T) {
	_, err := getPlaylistInfo("https://soundcloud.com/ilyanaazman/sets/best-of-mrrevillz")
	if err != nil {
		t.Error(err.Error())
		return
	}
}
