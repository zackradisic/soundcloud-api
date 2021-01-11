package soundcloudapi_test

import (
	"net/http"
	"testing"
	"time"

	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func TestNew(t *testing.T) {
	_, err := soundcloudapi.New(soundcloudapi.APIOptions{
		HTTPClient: &http.Client{
			Timeout: time.Second * 20,
		},
	})
	if err != nil {
		t.Errorf("failed to create new API: %+v\n", err)
	}
}
