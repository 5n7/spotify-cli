package spotify

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

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

func MustString(str string, err error) string {
	if err != nil {
		panic(err)
	}
	return str
}

func (c *CLI) StatusFromFormat(format string) (string, error) {
	tmpl := template.Must(template.New("status").Parse(format))
	status := struct {
		Album       string
		AlbumArtist string
		Artist      string
		Playback    string
		Title       string
		URL         string
	}{
		MustString(c.StatusFromKind(KindAlbum)),
		MustString(c.StatusFromKind(KindAlbumArtist)),
		MustString(c.StatusFromKind(KindArtist)),
		MustString(c.StatusFromKind(KindPlayback)),
		MustString(c.StatusFromKind(KindTitle)),
		MustString(c.StatusFromKind(KindURL)),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, status); err != nil {
		return ", err", err
	}
	return buf.String(), nil
}

func (c *CLI) StatusFromKind(kind Kind) (string, error) {
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
