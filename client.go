package soundcloudapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type client struct {
	httpClient *http.Client
	clientID   string
}

type failedRequestError struct {
	status int
	errMsg string
}

const trackURL = "https://api-v2.soundcloud.com/tracks"
const resolveURL = "https://api-v2.soundcloud.com/resolve"

func (f *failedRequestError) Error() string {
	if f.errMsg == "" {
		return fmt.Sprintf("Request returned non 2xx status: %d", f.status)
	}

	return fmt.Sprintf("Request failed with status %d: %s", f.status, f.errMsg)
}

func newClient(clientID string) *client {
	return &client{
		httpClient: http.DefaultClient,
		clientID:   clientID,
	}
}

func (c *client) makeRequest(method, url string, jsonBody interface{}) ([]byte, error) {
	var jsonBytes []byte
	var err error

	if jsonBody != nil {
		jsonBytes, err = json.Marshal(jsonBody)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to marshal json body")
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to make http request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		if data, err := ioutil.ReadAll(res.Body); err == nil {
			return nil, &failedRequestError{status: res.StatusCode, errMsg: string(data)}
		}
		return nil, &failedRequestError{status: res.StatusCode}
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		if data, err := ioutil.ReadAll(res.Body); err == nil {
			return nil, &failedRequestError{status: res.StatusCode, errMsg: string(data)}
		}
		return nil, &failedRequestError{status: res.StatusCode}
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return data, nil
	}

	return data, nil
}

func (c *client) buildURL(base string, clientID bool, query ...string) (string, error) {
	if len(query)%2 != 0 {
		return "", fmt.Errorf("Invalid query: URL (%s) Query: (%s)", base, strings.Join(query, ","))
	}

	u, err := url.Parse(string(base))
	if err != nil {
		return "", err
	}
	q := u.Query()

	for i := 0; i < len(query); i += 2 {
		q.Add(query[i], query[i+1])
	}

	if clientID {
		q.Add("client_id", c.clientID)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

// GetTrackInfoOptions can contain URL of the track or the ID of the track
type GetTrackInfoOptions struct {
	URL string
	ID  []int64
}

func (c *client) getTrackInfo(options GetTrackInfoOptions) ([]Track, error) {
	var u string
	var data []byte
	var err error

	var trackInfo []Track
	if options.ID != nil && len(options.ID) > 0 {
		ids := []string{}
		for _, id := range options.ID {
			ids = append(ids, strconv.FormatInt(id, 10))
		}
		u, err = c.buildURL(trackURL, true, "ids", strings.Join(ids, ","))
		if err != nil {
			return nil, errors.Wrap(err, "Failed to build URL for getTrackInfo()")
		}

		data, err = c.makeRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}
	} else if options.URL != "" {
		// TO-DO: Validate the URL
		data, err = c.resolve(options.URL)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &trackInfo)
	} else {
		return nil, errors.New("Invalid options. URL or ID must be provided")
	}

	err = json.Unmarshal(data, &trackInfo)

	if err != nil {
		return nil, errors.Wrap(err, "JSON is not valid track info")
	}

	if options.ID != nil && len(options.ID) > 0 {
		c.sortTrackInfo(options.ID, trackInfo)
	}

	return trackInfo, nil
}

func (c *client) sortTrackInfo(ids []int64, tracks []Track) {
	// Bubble Sort for now. Maybe switch to a more efficient sorting algorithm later??
	//
	// Because the API request in getTrackInfo is limited to 50 tracks at once
	// time complexity will always be <= O(50^2)
	for j, id := range ids {

		if tracks[j].ID != id {
			for k := 0; k < len(tracks); k++ {
				if tracks[k].ID == id {
					temp := tracks[j]
					tracks[j] = tracks[k]
					tracks[k] = temp
				}
			}
		}
	}
}

func (c *client) getPlaylistInfo(url string) (Playlist, error) {
	playlist := Playlist{}
	u, err := c.buildURL(resolveURL, true, "url", url)
	if err != nil {
		return playlist, errors.Wrap(err, "Failed to build URL for getPlaylistInfo")
	}

	data, err := c.makeRequest("GET", u, nil)
	if err != nil {
		return playlist, err
	}

	err = json.Unmarshal(data, &playlist)

	if err != nil {
		return playlist, errors.Wrap(err, "Returned JSON is not valid track info")
	}

	if playlist.TrackCount > 5 {
		ids := make([]int64, playlist.TrackCount-5)

		count := 0
		for _, track := range playlist.Tracks[5:] {
			ids[count] = track.ID
			count++
		}

		playlist.Tracks = playlist.Tracks[:5]

		if len(ids) > 50 {
			// The SoundCloud API limits querying tracks to 50 at a time.
			//
			// Split the requests.

			temp := make([]Track, len(ids))
			playlist.Tracks = append(playlist.Tracks, temp...)

			workers := len(ids) / 50

			type result struct {
				startIndex int
				trackInfo  []Track
			}

			errChan := make(chan error)
			resultsChan := make(chan result)
			for i := 0; i <= workers; i++ {
				start := i * 50
				end := start + 50
				if i == workers {
					end = start + (len(ids) % 50)
				}
				go func() {
					trackInfo, err := c.getTrackInfo(GetTrackInfoOptions{
						ID: ids[start:end],
					})

					if err != nil {
						errChan <- err
						return
					}

					resultsChan <- result{
						startIndex: start,
						trackInfo:  trackInfo,
					}
				}()
			}

			completeCount := -1

			for {
				select {
				case err = <-errChan:
					if err != nil {
						return playlist, errors.Wrap(err, "Failed to retreive playlist tracks")
					}
				case r := <-resultsChan:
					completeCount++

					for i, track := range r.trackInfo {
						playlist.Tracks[r.startIndex+i+5] = track
					}

					if completeCount == workers {
						break
					}
				}

				if completeCount == workers {
					break
				}
			}

		} else {
			trackInfo, err := c.getTrackInfo(GetTrackInfoOptions{
				ID: ids,
			})

			if err != nil {
				return playlist, errors.Wrap(err, "Failed to retrieve track information for playlist")
			}

			for _, track := range trackInfo {
				playlist.Tracks = append(playlist.Tracks, track)
			}
		}
	}

	data, err = json.Marshal(playlist)

	return playlist, nil
}

func (c *client) resolve(url string) ([]byte, error) {

	u, err := c.buildURL(resolveURL, true, "url", url)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to build URL for resolve()")
	}

	data, err := c.makeRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}
