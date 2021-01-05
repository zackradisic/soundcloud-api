package soundcloudapi_test

import (
	"testing"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func TestStripMobilePrefix(t *testing.T) {
	raw := "https://m.soundcloud.com/taliya-jenkins/double-cheese-burger-hold-the?ref=clipboard&p=i&c=0"
	expected := "https://soundcloud.com/taliya-jenkins/double-cheese-burger-hold-the?ref=clipboard&p=i&c=0"

	result := soundcloudapi.StripMobilePrefix(raw)

	if expected != result {
		t.Errorf("Expected: (%s), Received: (%s)\n", expected, result)
	}
}
