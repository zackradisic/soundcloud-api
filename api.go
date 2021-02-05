package soundcloudapi

// Track represents the JSON response of a track's info
type Track struct {
	Kind              string `json:"kind"`
	MonetizationModel string `json:"monetization_model"`
	ID                int64  `json:"id"`
	Policy            string `json:"polic"`
	CommentCount      int64  `json:"comment_count"`
	FullDurationMS    int64  `json:"full_duration"`
	Downloadable      bool   `json:"downloadable"`
	HasDownloadsLeft  bool   `json:"has_downloads_left"`
	CreatedAt         string `json:"created_at"`
	Description       string `json:"description"`
	Media             Media  `json:"media"`
	Title             string `json:"title"`
	DurationMS        int64  `json:"duration"`
	ArtworkURL        string `json:"artwork_url"`
	Public            bool   `json:"public"`
	Streamable        bool   `json:"streamable"`
	TagList           string `json:"tag_list"`
	Genre             string `json:"genre"`
	RepostsCount      int64  `json:"reposts_count"`
	LabelName         string `json:"label_name"`
	LastModified      string `json:"last_modified"`
	Commentable       bool   `json:"commentable"`
	URI               string `json:"uri"`
	DownloadCount     int64  `json:"download_count"`
	LikesCount        int64  `json:"likes_count"`
	DisplayDate       string `json:"display_date"`
	UserID            int64  `json:"user_id"`
	WaveformURL       string `json:"waveform_url"`
	Permalink         string `json:"permalink"`
	PermalinkURL      string `json:"permalink_url"`
	PlaybackCount     int64  `json:"playback_count"`
	SecretToken       string `json:"secret_token"`
	User              User   `json:"user"`
}

// Media contains an array of transcoding for a track
type Media struct {
	Transcodings []Transcoding `json:"transcodings"`
}

// Transcoding contains information about the transcoding of a track
type Transcoding struct {
	URL     string            `json:"url"`
	Preset  string            `json:"preset"`
	Snipped bool              `json:"snipped"`
	Format  TranscodingFormat `json:"format"`
}

// TranscodingFormat contains the protocol by which the track is delivered ("progressive" or "HLS"), and
// the mime type of the track
type TranscodingFormat struct {
	Protocol string `json:"protocol"`
	MimeType string `json:"mime_type"`
}

// Playlist represents the JSON response of a playlist
type Playlist struct {
	ArtworkURL     string  `json:"artwork_url"`
	CreatedAt      string  `json:"created_at"`
	Description    string  `json:"description"`
	DurationMS     int64   `json:"duration"`
	EmbeddeableBy  string  `json:"embeddable_by"`
	Genre          string  `json:"genre"`
	ID             int64   `json:"id"`
	Kind           string  `json:"kind"`
	LabelName      string  `json:"label_name"`
	LastModified   string  `json:"last_modified"`
	License        string  `json:"license"`
	LikesCount     int     `json:"likes_count"`
	ManagedByFeeds bool    `json:"managed_by_feeds"`
	Permalink      string  `json:"permalink"`
	PermalinkURL   string  `json:"permalink_url"`
	Public         bool    `json:"public"`
	SecretToken    string  `json:"secret_token"`
	Sharing        string  `json:"private"`
	TagList        string  `json:"tag_list"`
	Title          string  `json:"title"`
	URI            string  `json:"uri"`
	UserID         int64   `json:"user_id"`
	SetType        string  `json:"set_type"`
	IsAlbum        bool    `json:"is_album"`
	PublishedAt    string  `json:"published_at"`
	DisplayDate    string  `json:"display_date"`
	User           User    `json:"user"`
	Tracks         []Track `json:"tracks"`
	TrackCount     int     `json:"track_count"`
}

// User represents the JSON payload for user data
type User struct {
	ID              int64  `json:"id"`
	AvatarURL       string `json:"avatar_url"`
	City            string `json:"city"`
	CommentsCount   int64  `json:"comments_count"`
	CountryCode     string `json:"country_code"`
	CreatedAt       string `json:"created_at"`
	Description     string `json:"description"`
	FollowersCount  int64  `json:"followers_count"`
	FollowingsCount int64  `json:"followings_count"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PermalinkURL    string `json:"permalink_url"`
	URI             string `json:"uri"`
	Username        string `json:"username"`
	Kind            string `json:"kind"`
	Likes           int    `json:"likes_count"`
	PlaylistLikes   int    `json:"playlist_likes_count"`
	Verified        bool   `json:"verified"`
}

// MediaURLResponse is the JSON response of retrieving media information of a track
type MediaURLResponse struct {
	URL string `json:"url"`
}

// DownloadURLResponse is the JSON respose of retrieving media information of a publicly downloadable track
type DownloadURLResponse struct {
	URL string `json:"redirectUri"`
}

// PaginatedQuery is the JSON response for a paginated query
type PaginatedQuery struct {
	Collection   []map[string]interface{} `json:"collection"`
	TotalResults int                      `json:"total_results"`
	NextHref     string                   `json:"next_href"`
	QueryURN     string                   `json:"query_urn"`
}

// Like is the JSON response for a like
type Like struct {
	CreatedAt string   `json:"created_at"`
	Kind      string   `json:"kind"`
	Track     Track    `json:"track"`
	Playlist  Playlist `json:"playlist"`
}
