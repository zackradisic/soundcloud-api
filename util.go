package soundcloudapi

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const urlRegexp = `^https?:\/\/(soundcloud\.com)\/(.*)$`

var urlRegex = regexp.MustCompile(urlRegexp)

// IsURL returns true if the provided url is a valid SoundCloud URL
func IsURL(url string) bool {
	return len(urlRegex.FindAllString(url, -1)) > 0
}

// IsPlaylist retuns true if the provided url is a valid SoundCloud playlist URL
func IsPlaylist(u string) bool {
	if !IsURL(u) {
		return false
	}

	uObj, err := url.Parse(u)
	if err != nil {
		return false
	}

	return strings.Contains(uObj.Path, "/sets/")
}

func sliceContains(slice []int64, x int64) bool {
	for _, i := range slice {
		if i == x {
			return true
		}
	}

	return false
}

func deleteEmptyTracks(slice []Track) []Track {
	newTracks := []Track{}
	for _, t := range slice {
		if t.ID != 0 {
			newTracks = append(newTracks, t)
		}
	}

	return newTracks
}

// ExtractIDFromPersonalizedTrackURL extracts the track ID from a personalized track URL, returns -1
// if no track ID can be extracted
func ExtractIDFromPersonalizedTrackURL(url string) int64 {
	if !strings.Contains(url, "https://soundcloud.com/discover/sets/personalized-tracks::") {
		return -1
	}

	split := strings.Split(url, ":")
	if len(split) < 5 {
		return -1
	}

	id, err := strconv.ParseInt(split[4], 10, 64)
	if err != nil {
		return -1
	}

	return id
}
