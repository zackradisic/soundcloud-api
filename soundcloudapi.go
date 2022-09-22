package soundcloudapi

import (
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// API is a wrapper for the SoundCloud private API used internally for soundcloud.com
type API struct {
	client              *client
	StripMobilePrefix   bool
	ConvertFirebaseURLs bool
}

// APIOptions are the options for creating an API struct
type APIOptions struct {
	ClientID            string       // optional and a new one will be fetched if not provided
	HTTPClient          *http.Client // the HTTP client to make requests with
	StripMobilePrefix   bool         // whether or not to convert mobile URLs to regular URLs
	ConvertFirebaseURLs bool         // whether or not to convert SoundCloud firebase URLs to regular URLs
}

// New returns a pointer to a new SoundCloud API struct.
func New(options APIOptions) (*API, error) {

	if options.ClientID == "" {
		var err error
		options.ClientID, err = FetchClientID()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to initiaze SounCloudAPI")
		}
	}

	if options.HTTPClient == nil {
		options.HTTPClient = http.DefaultClient
	}

	return &API{
		client:              newClient(options.ClientID, options.HTTPClient),
		StripMobilePrefix:   options.StripMobilePrefix,
		ConvertFirebaseURLs: options.ConvertFirebaseURLs,
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
		url, err := sc.prepareURL(options.URL)
		if err != nil {
			return nil, err
		}
		options.URL = url
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
	url, err := sc.prepareURL(transcoding.URL)
	if err != nil {
		return err
	}
	u, err := sc.client.getMediaURL(url)
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
	url, err := sc.prepareURL(options.ProfileURL)
	if err != nil {
		return nil, err
	}
	options.ProfileURL = url
	return sc.client.getLikes(options)
}

// Search returns a PaginatedQuery for searching a specific query
func (sc *API) Search(options SearchOptions) (*PaginatedQuery, error) {
	return sc.client.search(options)
}

// GetUser returns a User
func (sc *API) GetUser(options GetUserOptions) (User, error) {
	url, err := sc.prepareURL(options.ProfileURL)
	if err != nil {
		return User{}, err
	}
	options.ProfileURL = url
	return sc.client.getUser(options)
}

// GetDownloadURL retuns the URL to download a track. This is useful if you want to implement your own
// downloading algorithm.
// If the track has a publicly available download link, that link will be preferred and the streamType parameter will be ignored.
// streamType can be either "hls" or "progressive", defaults to "progressive"
func (sc *API) GetDownloadURL(url string, streamType string) (string, error) {
	url, err := sc.prepareURL(url)
	if err != nil {
		return "", err
	}
	streamType = strings.ToLower(streamType)
	if streamType == "" {
		streamType = "progressive"
	}

	if IsURL(url, false, false) && !IsPlaylistURL(url) {
		info, err := sc.client.getTrackInfo(GetTrackInfoOptions{
			URL: url,
		})

		if err != nil {
			return "", err
		}

		if len(info) == 0 {
			return "", errors.New("Could not find a track with that URL")
		}

		if info[0].Downloadable && info[0].HasDownloadsLeft {
			downloadURL, err := sc.client.getDownloadURL(info[0].ID)
			if err != nil {
				return "", err
			}
			return downloadURL, nil
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

		mediaURL, err := sc.client.getMediaURL(info[0].Media.Transcodings[0].URL)
		if err != nil {
			return "", err
		}
		return mediaURL, nil
	}
	return "", errors.New("URL is not a track URL")
}

func (sc *API) prepareURL(url string) (string, error) {
	if sc.StripMobilePrefix {
		if IsMobileURL(url) {
			url = StripMobilePrefix(url)
		}
	}

	if sc.ConvertFirebaseURLs {
		if IsFirebaseURL(url) {
			var err error
			url, err = ConvertFirebaseLink(url)
			if err != nil {
				return "", errors.Wrap(err, "Failed to convert Firebase URL")
			}
		}
	}

	if IsNewMobileURL(url) {
		var err error
		url, err = sc.ConvertNewMobileURL(url)
		if err != nil {
			return "", errors.Wrap(err, "failed to convert new mobile url")
		}
	}

	return url, nil
}

func (sc *API) ConvertNewMobileURL(url string) (string, error) {
	client := new(http.Client)
	type urlResp struct {
		url *string
		err error
	}
	urlChan := make(chan urlResp, 1)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		u := req.URL.String()

		if IsURL(u, false, false) {
			urlChan <- urlResp{url: &u, err: nil}
		}

		return nil
	}

	_, err := client.Get(url)
	select {
	case urlR := <-urlChan:
		if urlR.url == nil {
			return "", errors.New("unable to retrieve redirect url for new mobile url")
		}
		return *urlR.url, nil
	default:
		if err != nil {
			return "", errors.Wrap(err, "failed to get redirect url")
		}
		return "", errors.New("new mobile url is supposed to have redirects")
	}

}

// IsURL is a shorthand for IsURL(url, sc.StripMobilePrefix, sc.ConvertFirebaseURLs)
func (sc *API) IsURL(url string) bool {
	return IsURL(url, sc.StripMobilePrefix, sc.ConvertFirebaseURLs)
}
