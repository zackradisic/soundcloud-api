![soundcloud-api](https://socialify.git.ci/zackradisic/soundcloud-api/image?description=1&language=1&owner=1&pattern=Plus&stargazers=1&theme=Dark)

[![GoDoc](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go)](https://pkg.go.dev/github.com/zackradisic/soundcloud-api)

### SoundCloud's v2 API reverse engineered for Go.


# Notice
The SoundCloud api-v2 is an [undocumented, internal](https://stackoverflow.com/questions/29253633/soundcloud-is-api-v2-allowed-to-be-used-and-is-there-documentation-on-it) API used by the web app at https://soundcloud.com. 

SoundCloud is currently [not](https://docs.google.com/forms/d/e/1FAIpQLSfNxc82RJuzC0DnISat7n4H-G7IsPQIdaMpe202iiHZEoso9w/closedform) allowing developers to register for applications, and using undocumented APIs is apparently breaking SoundCloud's [ToS](https://twitter.com/SoundCloudDev/status/639017606264016896), use this at your own risk.

# Quick Start

```go
// You can specify certain options here if you want
sc, err := soundcloudapi.New(soundcloudapi.APIOptions{}) 

if err != nil {
    log.Fatal(err.Error())
}

track, err := sc.GetTrackInfo(soundcloudapi.GetTrackInfoOptions{
    URL: "https://soundcloud.com/track/infsdfo",
})
```

See the [docs](https://pkg.go.dev/github.com/zackradisic/soundcloud-api) for more reference.

# Error Handling
If an error is returned from SoundCloud's API, it will take the form of the FailedRequestError struct. You can use type
assertions to access the status code or JSON error msg for your use case. Ex:

```go
tracks, err := sc.GetTrackInfo(soundcloudapi.GetTrackInfoOptions{
    URL: "https://soundcloud.com/asdkfjhalsdhfl",
})

if failedRequest, ok := err.(*soundcloudapi.FailedRequestError); ok {
    if failedRequest.Status == 404 {
        fmt.Println("Could not find that track")
    }
    return
}
```

# Paginated Queries
Functions like [`sc.Search()`](https://pkg.go.dev/github.com/zackradisic/soundcloud-api@v0.1.0#API.Search) or [`sc.GetLikes()`](https://pkg.go.dev/github.com/zackradisic/soundcloud-api@v0.1.0#API.GetLikes) return a [PaginatedQuery](https://pkg.go.dev/github.com/zackradisic/soundcloud-api@v0.1.0#PaginatedQuery). PaginatedQuery.Collection contains the JSON of the items that matched the query,
represented as a `map[string]interface{}`. You can use the provided functions to get the items in the form you want:

```go
paginatedQuery, _ := sc.Search(soundcloudapi.SearchOptions{
    Query: "childish gambino"
})

tracks, _ := paginatedQuery.GetTracks() // Get the tracks of the response
playlists, _ := paginatedQuery.GetPlaylists() // Get the playlists of the response
likes, _ := paginatedQuery.GetLikes() // Get the likes of the response
```