package cmd

import (
	"fmt"

	"github.com/skmatz/spotify-cli"
	"github.com/spf13/cobra"
)

var (
	format string
	kind   string
)

func runFormat(cli *spotify.CLI, format string) error {
	f, err := cli.StatusFromFormat(format)
	if err != nil {
		return err
	}
	fmt.Println(f)
	return nil
}

func runKind(cli *spotify.CLI, kind string) error {
	switch kind {
	case "album":
		album, err := cli.StatusFromKind(spotify.KindAlbum)
		if err != nil {
			return err
		}
		fmt.Println(album)
	case "album-artist":
		albumArtist, err := cli.StatusFromKind(spotify.KindAlbumArtist)
		if err != nil {
			return err
		}
		fmt.Println(albumArtist)
	case "artist":
		artist, err := cli.StatusFromKind(spotify.KindArtist)
		if err != nil {
			return err
		}
		fmt.Println(artist)
	case "playback":
		playback, err := cli.StatusFromKind(spotify.KindPlayback)
		if err != nil {
			return err
		}
		fmt.Println(playback)
	case "title":
		title, err := cli.StatusFromKind(spotify.KindTitle)
		if err != nil {
			return err
		}
		fmt.Println(title)
	case "url":
		url, err := cli.StatusFromKind(spotify.KindURL)
		if err != nil {
			return err
		}
		fmt.Println(url)
	default:
		return fmt.Errorf("invalid kind: %s", kind)
	}
	return nil
}

func runStatus(cmd *cobra.Command, args []string) error {
	cli, err := spotify.New()
	if err != nil {
		return err
	}

	if format != "" {
		return runFormat(cli, format)
	}

	if kind != "" {
		return runKind(cli, kind)
	}
	return fmt.Errorf("specify either format or kind")
}

var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"s"},
	Short:   "Get Spotify status via DBus",
	Long:    "Get Spotify status via DBus.",
	RunE:    runStatus,
}

func init() { //nolint:gochecknoinits
	statusCmd.Flags().StringVarP(&format, "format", "f", "",
		"format ({{ .Album }}|{{ .AlbumArtist }}|{{ .Artist }}|{{ .Playback }}|{{ .Title }}|{{ .URL }})")
	statusCmd.Flags().StringVarP(&kind, "kind", "k", "", "kind (album|album-artist|artist|plackback|title|url)")

	rootCmd.AddCommand(statusCmd)
}
