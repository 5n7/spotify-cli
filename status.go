package spotify

import (
	"fmt"
	"strings"

	"github.com/dawidd6/go-spotify-dbus"
)

type Kind int

const (
	KindAlbum Kind = iota
	KindAlbumArtist
	KindArtist
	KindPlayback
	KindTitle
	KindURL
)

var ListSeparator string = ", "

func (c *CLI) Status(kind Kind) (string, error) {
	metadata, err := spotify.GetMetadata(c.conn)
	if err != nil {
		return "", err
	}

	playback, err := spotify.GetPlaybackStatus(c.conn)
	if err != nil {
		return "", err
	}

	switch kind {
	case KindAlbum:
		return metadata.Album, nil
	case KindAlbumArtist:
		return strings.Join(metadata.AlbumArtist, ListSeparator), nil
	case KindArtist:
		return strings.Join(metadata.Artist, ListSeparator), nil
	case KindPlayback:
		return string(playback), nil
	case KindTitle:
		return metadata.Title, nil
	case KindURL:
		return metadata.URL, nil
	default:
		return "", fmt.Errorf("the kind is not supported")
	}
}
