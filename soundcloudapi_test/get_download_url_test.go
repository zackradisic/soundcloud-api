package soundcloudapi_test

import (
	"strings"
	"testing"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func TestGetDownloadURL(t *testing.T) {
	dlURL, err := api.GetDownloadURL("https://soundcloud.com/taliya-jenkins/double-cheese-burger-hold-the", "hls")
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !strings.Contains(dlURL, "sndcdn.com/") {
		t.Errorf("Invalid download URL returned, received: (%s)", dlURL)
	}
}

func TestGetDownloadURLPublic(t *testing.T) {
	// This track has a public download URL link
	trackInfo, err := api.GetTrackInfo(soundcloudapi.GetTrackInfoOptions{
		URL: "https://soundcloud.com/moccioso/01_wav?in=moccioso/sets/download-converter-test",
	})
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !trackInfo[0].Downloadable {
		t.Error("Track changed, update the URL")
		return
	}

	dlURL, err := api.GetDownloadURL("https://soundcloud.com/moccioso/01_wav?in=moccioso/sets/download-converter-test", "")
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !strings.Contains(dlURL, "sndcdn.com/") {
		t.Errorf("Invalid download URL returned, received: (%s)", dlURL)
	}
}
