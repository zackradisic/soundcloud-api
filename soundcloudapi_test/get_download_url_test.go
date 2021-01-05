package soundcloudapi_test

import (
	"strings"
	"testing"
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
