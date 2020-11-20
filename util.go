package soundcloudapi

import (
	"regexp"
)

const urlRegexp = `^https?:\/\/(soundcloud\.com)\/(.*)$`

var urlRegex = regexp.MustCompile(urlRegexp)

// IsURL returns true if the provided url is a valid SoundCloud URL
func IsURL(url string) bool {
	return len(urlRegex.FindAllString(url, -1)) > 0
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
