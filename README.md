![soundcloud-api](https://socialify.git.ci/zackradisic/soundcloud-api/image?description=1&language=1&owner=1&pattern=Plus&stargazers=1&theme=Dark)

[![GoDoc](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go)](https://pkg.go.dev/github.com/zackradisic/soundcloud-api)

### SoundCloud's v2 API reverse engineered for Go.


# Notice
The SoundCloud api-v2 is an [undocumented, internal](https://stackoverflow.com/questions/29253633/soundcloud-is-api-v2-allowed-to-be-used-and-is-there-documentation-on-it) API used by the web app at https://soundcloud.com. 

SoundCloud is currently [not](https://docs.google.com/forms/d/e/1FAIpQLSfNxc82RJuzC0DnISat7n4H-G7IsPQIdaMpe202iiHZEoso9w/closedform) allowing developers to register for applications, and using undocumented APIs is apparently breaking SoundCloud's [ToS](https://twitter.com/SoundCloudDev/status/639017606264016896), use this at your own risk.

# Quick Start

```go
// You can pass in a client ID if you want to, 
// if not the package will fetch one for you
sc, err := soundcloudapi.NewSoundCloudAPI("") 

if err != nil {
    log.Fatal(err.Error())
}

track, err := sc.GetTrackInfo(soundcloudapi.GetTrackInfoOptions{
		URL: "https://soundcloud.com/track/infsdfo",
})
```

See the [docs](https://pkg.go.dev/github.com/zackradisic/soundcloud-api) for more reference.
