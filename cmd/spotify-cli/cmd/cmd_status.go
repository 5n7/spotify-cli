package cmd

import (
	"fmt"

	"github.com/skmatz/spotify-cli"
	"github.com/spf13/cobra"
)

var kind string

func runStatus(cmd *cobra.Command, args []string) error {
	cli, err := spotify.New()
	if err != nil {
		return err
	}

	switch kind {
	case "album":
		album, err := cli.Status(spotify.KindAlbum)
		if err != nil {
			return err
		}
		fmt.Println(album)
	case "album-artist":
		albumArtist, err := cli.Status(spotify.KindAlbumArtist)
		if err != nil {
			return err
		}
		fmt.Println(albumArtist)
	case "artist":
		artist, err := cli.Status(spotify.KindArtist)
		if err != nil {
			return err
		}
		fmt.Println(artist)
	case "playback":
		playback, err := cli.Status(spotify.KindPlayback)
		if err != nil {
			return err
		}
		fmt.Println(playback)
	case "title":
		title, err := cli.Status(spotify.KindTitle)
		if err != nil {
			return err
		}
		fmt.Println(title)
	case "url":
		url, err := cli.Status(spotify.KindURL)
		if err != nil {
			return err
		}
		fmt.Println(url)
	default:
		return fmt.Errorf("invalid kind: %s", kind)
	}
	return nil
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get Spotify status via DBus",
	Long:  "Get Spotify status via DBus.",
	RunE:  runStatus,
}

func init() { //nolint:gochecknoinits
	statusCmd.Flags().StringVar(&kind, "kind", "", "kind (album|album-artist|artist|plackback|title|url)")

	rootCmd.AddCommand(statusCmd)
}
