![soundcloud-api](https://socialify.git.ci/zackradisic/soundcloud-api/image?description=1&language=1&owner=1&pattern=Plus&stargazers=1&theme=Dark)

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
