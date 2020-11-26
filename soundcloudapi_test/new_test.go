package soundcloudapi_test

import (
	"testing"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func TestNew(t *testing.T) {
	_, err := soundcloudapi.New("")
	if err != nil {
		t.Errorf("failed to create new API: %+v\n", err)
	}
}
