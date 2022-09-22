package soundcloudapi

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var firebaseRegex = regexp.MustCompile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._+~#=]{1,500}\\.[a-zA-Z0-9()]{1,500}\\b([-a-zA-Z0-9()@:%_+.~#?&//\\\\=]*)")

var urlRegex = regexp.MustCompile(`(?m)^https?:\/\/(soundcloud\.com)\/(.*)$`)

var firebaseURLRegex = regexp.MustCompile(`(?m)^https?:\/\/(soundcloud\.app\.goo\.gl)\/(.*)$`)

var mobileURLRegex = regexp.MustCompile(`(?m)^https?:\/\/(m\.soundcloud\.com)\/(.*)$`)

var unicodeRegex = regexp.MustCompile(`(?i)\\u([\d\w]{4})`)

// IsURL returns true if the provided url is a valid SoundCloud URL
func IsURL(url string, testMobile, testFirebase bool) bool {
	success := false
	if testMobile {
		success = IsMobileURL(url)
	}

	if testFirebase && !success {
		success = IsFirebaseURL(url)
	}

	if !success {
		success = len(urlRegex.FindAllString(url, -1)) > 0
	}

	return success
}

// StripMobilePrefix removes the prefix for mobile urls. Returns the same string if an error parsing the URL occurs
func StripMobilePrefix(u string) string {
	if !strings.Contains(u, "m.soundcloud.com") {
		return u
	}
	_url, err := url.Parse(u)
	if err != nil {
		return u
	}
	_url.Host = "soundcloud.com"
	return _url.String()
}

// IsFirebaseURL returns true if the url is a SoundCloud Firebase url (has the following form: https://soundcloud.app.goo.gl/xxxxxxxx)
func IsFirebaseURL(u string) bool {
	return len(firebaseURLRegex.FindAllString(u, -1)) > 0
}

// IsMobileURL returns true if the url is a SoundCloud Firebase url (has the following form: https://m.soundcloud.com/xxxxxx)
func IsMobileURL(u string) bool {
	return len(mobileURLRegex.FindAllString(u, -1)) > 0
}

func IsNewMobileURL(u string) bool {
	return strings.Index(u, "https://on.soundcloud.com/") == 0
}

func replaceUnicodeChars(str string) (string, error) {
	for _, match := range unicodeRegex.FindAllString(str, -1) {
		s, err := strconv.Unquote("'" + match + "'")
		if err != nil {
			return "", err
		}
		str = strings.Replace(str, match, s, -1)
	}

	return str, nil
}

// ConvertFirebaseLink converts a link of the form (https://soundcloud.app.goo.gl/xxxxxxxx) to a regular
// SoundCloud link.
func ConvertFirebaseLink(u string) (string, error) {
	_url, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	q := _url.Query()
	q.Set("d", "1")
	_url.RawQuery = q.Encode()

	res, err := http.Get(_url.String())
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	matches := firebaseRegex.FindAllString(string(data), -1)

	for _, match := range matches {
		if IsURL(match, false, false) {
			str, err := replaceUnicodeChars(match)
			if err != nil {
				return "", err
			}
			return str, nil
		}
	}

	return "", errors.New("Could not find regular SoundCloud URL from the URL provided")
}

// IsPlaylistURL retuns true if the provided url is a valid SoundCloud playlist URL
func IsPlaylistURL(u string) bool {
	if !IsURL(u, false, false) {
		return false
	}

	if IsPersonalizedTrackURL(u) {
		return false
	}

	uObj, err := url.Parse(u)
	if err != nil {
		return false
	}

	return strings.Contains(uObj.Path, "/sets/")
}

// IsSearchURL returns true  if the provided url is a valid search url
func IsSearchURL(url string) bool {
	return strings.Index(url, "https://soundcloud.com/search?") == 0
}

// IsPersonalizedTrackURL returns true if the provided url is a valid personalized track url. Ex/
// https://soundcloud.com/discover/sets/personalized-tracks::sam:335899198
func IsPersonalizedTrackURL(url string) bool {
	return strings.Contains(url, "https://soundcloud.com/discover/sets/personalized-tracks::")
}

// ExtractIDFromPersonalizedTrackURL extracts the track ID from a personalized track URL, returns -1
// if no track ID can be extracted
func ExtractIDFromPersonalizedTrackURL(url string) int64 {
	if !IsPersonalizedTrackURL(url) {
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
