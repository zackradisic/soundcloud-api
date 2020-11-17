package soundcloudapi

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// FetchClientID fetches a SoundCloud client ID.
// This algorithm is adapted from:
//     https://www.npmjs.com/package/soundcloud-key-fetch
func FetchClientID() (string, error) {
	// // // // // // // // // // // // // // // // // // // // // // // // // // // // //
	// 																					//
	// The basic notion of how this function works is that SoundCloud provides          //
	// a client ID so its web app can make API requests.								//
	//																					//
	// This client ID (along with other intialization data for the web app) is provided //
	// in a JavaScript file imported through a <script> tag in the HTML.				//
	//																					//
	// This function scrapes the HTML and tries to find the URL to that JS file,		//
	// and then scrapes the JS file to find the client ID.								//
	//																					//
	// // // // // // // // // // // // // // // // // // // // // // // // // // // // //

	resp, err := http.Get("https://soundcloud.com")
	if err != nil {
		return "", errors.Wrap(err, "Failed to fetch SoundCloud Client ID")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read body while fetching SoundCloud Client ID")
	}

	bodyString := string(body)
	split := strings.Split(bodyString, `<script crossorigin src="`)
	urls := []string{}

	for _, raw := range split {
		u := strings.Replace(raw, `"></script>`, "", 1)
		u = strings.Split(u, "\n")[0]
		if string([]rune(u)[0:31]) == "https://a-v2.sndcdn.com/assets/" {
			urls = append(urls, u)
		}
	}

	resp, err = http.Get(urls[len(urls)-1])
	if err != nil {
		return "", errors.Wrap(err, "Failed to fetch SoundCloud Client ID")
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read body while fetching SoundCloud Client ID")
	}

	bodyString = string(body)

	if strings.Contains(bodyString, `,client_id:"`) {
		clientID := strings.Split(bodyString, `,client_id:"`)[1]
		clientID = strings.Split(clientID, `"`)[0]
		return clientID, nil
	}

	return "", errors.New("Could not find a SoundCloud client ID")
}
