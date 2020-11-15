# soundcloud-api

A SoundCloud v2 API wrapper designed.

# Quick Start

```go
sc, err := soundcloudapi.NewSoundCloudAPI("")

if err != nil {
    log.Fatal(err.Error())
}

track, err := sc.GetTrackInfo(soundcloudapi.GetTrackInfoOptions{
		URL: "https://soundcloud.com/track/infsdfo",
})
```
