package soundcloudapi_test

import (
	"testing"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func TestConvertFirebaseLinkPrefix(t *testing.T) {
	raw := "https://soundcloud.app.goo.gl/z8snjNyHU8zMHH29A"
	expected := "https://soundcloud.com/taliya-jenkins/double-cheese-burger-hold-the?ref=clipboard&p=i&c=0"

	result, err := soundcloudapi.ConvertFirebaseLink(raw)

	if err != nil {
		t.Error(err.Error())
		return
	}

	if expected != result {
		t.Errorf("Expected: (%s), Received: (%s)\n", expected, result)
	}
}
