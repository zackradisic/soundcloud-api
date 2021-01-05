package soundcloudapi

import (
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// API is a wrapper for the SoundCloud private API used internally for soundcloud.com
type API struct {
	client *client
}

// New returns a pointer to a new SoundCloud API struct.
//
// clientID is optional and a new one will be fetched if not provided
func New(clientID string, client *http.Client) (*API, error) {
	if clientID == "" {
		var err error
		clientID, err = FetchClientID()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to initiaze SounCloudAPI")
		}
	}

	return &API{
		client: newClient(clientID, client),
	}, nil
}

// SetClientID sets the client ID
func (sc *API) SetClientID(clientID string) {
	sc.client.clientID = clientID
}

// ClientID returns the client ID
func (sc *API) ClientID() string {
	return sc.client.clientID
}

// GetTrackInfo returns the info for the track given tracks
//
// If URL is supplied, it will return the info for a single track given by that url.
// If an array of ids is supplied, it will return an array of track info.
//
// WARNING: Private tracks will not be fetched unless options.PlaylistID and options.PlaylistSecretToken
// are provided.
func (sc *API) GetTrackInfo(options GetTrackInfoOptions) ([]Track, error) {
	if options.URL != "" {
		options.URL = StripMobilePrefix(options.URL)
		id := ExtractIDFromPersonalizedTrackURL(options.URL)
		if id != -1 {
			return sc.client.getTrackInfo(GetTrackInfoOptions{ID: []int64{id}})
		}
	}
	return sc.client.getTrackInfo(options)
}

// GetPlaylistInfo returns the info for a playlist
func (sc *API) GetPlaylistInfo(url string) (Playlist, error) {
	return sc.client.getPlaylistInfo(StripMobilePrefix(url))
}

// DownloadTrack downloads the track specified by the given Transcoding's URL to dst
func (sc *API) DownloadTrack(transcoding Transcoding, dst io.Writer) error {
	u, err := sc.client.getMediaURL(StripMobilePrefix(transcoding.URL))
	if err != nil {
		return err
	}
	if strings.Contains(transcoding.URL, "progressive") {
		// Progressive download
		err = sc.client.downloadProgressive(u, dst)
	} else {
		// HLS download
		err = sc.client.downloadHLS(u, dst)
	}

	return err
}

// GetLikes returns a PaginatedQuery with the Collection field member as a list of tracks
func (sc *API) GetLikes(options GetLikesOptions) (*PaginatedQuery, error) {
	options.ProfileURL = StripMobilePrefix(options.ProfileURL)
	return sc.client.getLikes(options)
}

// Search returns a PaginatedQuery for searching a specific query
func (sc *API) Search(options SearchOptions) (*PaginatedQuery, error) {
	return sc.client.search(options)
}

// GetUser returns a User
func (sc *API) GetUser(options GetUserOptions) (User, error) {
	options.ProfileURL = StripMobilePrefix(options.ProfileURL)
	return sc.client.getUser(options)
}

// GetDownloadURL retuns the URL to download a track. This is useful if you want to implement your own
// downloading algorithm.
//
// streamType can be either "hls" or "progressive", defaults to "progressive"
func (sc *API) GetDownloadURL(url string, streamType string) (string, error) {
	url = StripMobilePrefix(url)
	streamType = strings.ToLower(streamType)
	if streamType == "" {
		streamType = "progressive"
	}

	if IsURL(url) && !IsPlaylistURL(url) {
		info, err := sc.client.getTrackInfo(GetTrackInfoOptions{
			URL: url,
		})

		if err != nil {
			return "", err
		}

		if len(info) == 0 {
			return "", errors.New("Could not find a track with that URL")
		}

		for _, transcoding := range info[0].Media.Transcodings {
			if strings.ToLower(transcoding.Format.Protocol) == streamType {
				mediaURL, err := sc.client.getMediaURL(transcoding.URL)
				if err != nil {
					return "", err
				}
				return mediaURL, nil
			}
		}
	} else {
		return "", errors.New("URL is not a track URL")
	}

	return "", errors.New("Could not find a download URL for that track")
}
