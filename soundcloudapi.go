package soundcloudapi

import (
	"io"
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
func New(clientID string) (*API, error) {
	if clientID == "" {
		var err error
		clientID, err = FetchClientID()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to initiaze SounCloudAPI")
		}
	}

	return &API{
		client: newClient(clientID),
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
		id := ExtractIDFromPersonalizedTrackURL(options.URL)
		if id != -1 {
			return sc.client.getTrackInfo(GetTrackInfoOptions{ID: []int64{id}})
		}
	}
	return sc.client.getTrackInfo(options)
}

// GetPlaylistInfo returns the info for a playlist
func (sc *API) GetPlaylistInfo(url string) (Playlist, error) {
	return sc.client.getPlaylistInfo(url)
}

// DownloadTrack downloads the track specified by the given Transcoding's URL to dst
func (sc *API) DownloadTrack(transcoding Transcoding, dst io.Writer) error {
	u, err := sc.client.getMediaURL(transcoding.URL)
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
func (sc *API) GetLikes(options GetLikesOptions) (*PaginatedLikeQuery, error) {
	return sc.client.getLikes(options)
}
