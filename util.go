package soundcloudapi

import "regexp"

const urlRegexp = `^https?:\/\/(soundcloud\.com)\/(.*)$`

var urlRegex = regexp.MustCompile(urlRegexp)

// IsURL returns true if the provided url is a valid SoundCloud URL
func IsURL(url string) bool {
	return len(urlRegex.FindAllString(url, -1)) > 0
}
