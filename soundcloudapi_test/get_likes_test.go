package soundcloudapi_test

import (
	"net/url"
	"os"
	"testing"

	"github.com/pkg/errors"
	soundcloudapi "github.com/zackradisic/soundcloud-api"
)

func TestGetLikes(t *testing.T) {
	limit := 10
	options := soundcloudapi.GetLikesOptions{
		ProfileURL: "https://soundcloud.com/jaiseanforever",
		Limit:      limit,
		Type:       "track",
	}

	likes, err := getLikes(options)
	if err != nil {
		t.Error(err.Error())
		return
	}

	for _, like := range likes {
		if like.Track.Kind != "track" {
			t.Errorf("Like is for the wrong type of resource: Expected (%s), received (%s)\n", "track", like.Track.Kind)
			return
		}
	}

	options.Type = "playlist"
	options.ProfileURL = "https://soundcloud.com/ibr"

	likes, err = getLikes(options)
	if err != nil {
		t.Error(err.Error())
		return
	}

	for _, like := range likes {
		if like.Playlist.Kind != "playlist" {
			t.Errorf("Like is for the wrong type of resource: Expected (%s), received (%s)\n", "playlist", like.Playlist.Kind)
			return
		}
	}

	options.Type = "track"
	options.ProfileURL = ""
	options.ID = 304506184

	likes, err = getLikes(options)
	if err != nil {
		t.Errorf("getLikes with ID failed: %s\n", err.Error())
		return
	}

}

func getLikes(options soundcloudapi.GetLikesOptions) ([]soundcloudapi.Like, error) {
	response, err := api.GetLikes(options)

	if err != nil {
		return nil, errors.Wrap(err, "GetLikes API endpoint failed")
	}

	if len(response.Collection) > options.Limit {
		return nil, errors.Errorf("Collection does not have the right max amount of items. Expected max (%d), received (%d)\n", options.Limit, len(response.Collection))
	}

	for _, item := range response.Collection {
		if val, ok := item["kind"]; ok {
			if kind, ok := val.(string); ok {
				if kind != "like" {
					return nil, errors.Errorf("Collection item has wrong value for 'kind' property. Expected (%s), received (%s)\n", "like", kind)
				}
			} else {
				return nil, errors.Errorf("Collection item has wrong type for 'kind' property. Expected (%s), received (%T)", "string", val)
			}
		} else {
			return nil, errors.New("Collection item has no 'kind' property")
		}
	}

	likes, err := response.GetLikes()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get likes from response")
	}

	if len(likes) != options.Limit {
		return nil, errors.Errorf("Amount of likes in collection does not match limit parameter. Length: %d", len(likes))
	}

	return likes, nil
}

func TestGetLikesBulk(t *testing.T) {
	// Skip this test on CI since it will github actions up
	if os.Getenv("CI") != "" {
		return
	}
	actualLimit := 2000
	options := soundcloudapi.GetLikesOptions{
		ProfileURL: "https://soundcloud.com/dasc2000",
		Limit:      1000,
		Type:       "track",
	}

	likesMap := map[string]struct{}{}

	i := 0
	for {
		likes, err := api.GetLikes(options)
		if err != nil {
			t.Error(err.Error())
			return
		}

		l, err := likes.GetLikes()
		if err != nil {
			t.Error(err.Error())
			return
		}

		for _, like := range l {
			if like.Track.Kind != "" {
				_, exists := likesMap[like.Track.Title]
				if !exists {
					likesMap[like.Track.Title] = struct{}{}
					i++
				}
			}
		}

		if likes.NextHref == "" {
			return
		}

		u, err := url.Parse(likes.NextHref)
		if err != nil {
			panic(err)
		}
		offset := u.Query().Get("offset")
		options.Offset = offset
		if i >= actualLimit {
			return
		}
	}
}
