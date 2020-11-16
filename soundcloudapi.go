package soundcloudapi

import (
	"github.com/pkg/errors"
)

// SoundCloudAPI is a wrapper for the SoundCloud private API used internally for soundcloud.com
type SoundCloudAPI struct {
	client *client
}

// NewSoundCloudAPI returns a pointer to a SoundCloud API struct.
//
// clientID is optional and a new one will be fetched if not provided
func NewSoundCloudAPI(clientID string) (*SoundCloudAPI, error) {
	if clientID == "" {
		var err error
		clientID, err = FetchClientID()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to initiaze SounCloudAPI")
		}
	}

	return &SoundCloudAPI{
		client: newClient(clientID),
	}, nil
}

// SetClientID sets the client ID
func (sc *SoundCloudAPI) SetClientID(clientID string) {
	sc.client.clientID = clientID
}

// ClientID returns the client ID
func (sc *SoundCloudAPI) ClientID() string {
	return sc.client.clientID
}

// GetTrackInfo returns the info for the track given tracks
//
// If URL is supplied, it will return the info for a single track given by that url.
// If an array of ids is supplied, it will return an array of track info
func (sc *SoundCloudAPI) GetTrackInfo(options GetTrackInfoOptions) ([]Track, error) {
	return sc.client.getTrackInfo(options)
}

// GetPlaylistInfo returns the info for a playlist
func (sc *SoundCloudAPI) GetPlaylistInfo(url string) (Playlist, error) {
	return sc.client.getPlaylistInfo(url)
}
