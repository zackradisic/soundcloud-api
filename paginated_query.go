package soundcloudapi

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// GetTracks returns any of the items in the PaginatedQuery's collection that match the Track struct type
func (pq *PaginatedQuery) GetTracks() ([]Track, error) {

	tracks := make([]Track, 0)

	for _, item := range pq.Collection {
		track := Track{}
		b, err := json.Marshal(item)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to marshal PaginatedQuery collection item")
		}

		err = json.Unmarshal(b, &track)
		if err != nil {
			continue
		}

		if track.Kind != "track" {
			continue
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
}

// GetPlaylists returns any of the items in the PaginatedQuery's collection that match the Playlist struct type
func (pq *PaginatedQuery) GetPlaylists() ([]Playlist, error) {
	playlists := make([]Playlist, 0)

	for _, item := range pq.Collection {
		playlist := Playlist{}
		b, err := json.Marshal(item)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to marshal PaginatedQuery collection item")
		}

		err = json.Unmarshal(b, &playlist)
		if err != nil {
			continue
		}

		if playlist.Kind != "playlist" {
			continue
		}

		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

// GetLikes returns any of the items in the PaginatedQuery's collection that match the Like struct type
func (pq *PaginatedQuery) GetLikes() ([]Like, error) {
	likes := make([]Like, 0)

	for _, item := range pq.Collection {
		like := Like{}
		b, err := json.Marshal(item)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to marshal PaginatedQuery collection item")
		}

		err = json.Unmarshal(b, &like)
		if err != nil {
			continue
		}

		if like.Kind != "like" {
			continue
		}

		likes = append(likes, like)
	}

	return likes, nil
}
