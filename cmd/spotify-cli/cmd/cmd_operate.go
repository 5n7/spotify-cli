package cmd

import (
	"fmt"

	"github.com/skmatz/spotify-cli"
	"github.com/spf13/cobra"
)

var action string

func runOperate(cmd *cobra.Command, args []string) error {
	cli, err := spotify.New()
	if err != nil {
		return err
	}

	switch action {
	case "next":
		return cli.Operate(spotify.ActionNext)
	case "pause":
		return cli.Operate(spotify.ActionPause)
	case "play":
		return cli.Operate(spotify.ActionPlay)
	case "play-pause":
		return cli.Operate(spotify.ActionPlayPause)
	case "previous":
		return cli.Operate(spotify.ActionPrevious)
	default:
		return fmt.Errorf("invalid action: %s", action)
	}
}

var operateCmd = &cobra.Command{
	Use:   "operate",
	Short: "Operate Spotify via DBus",
	Long:  "Operate Spotify via DBus.",
	RunE:  runOperate,
}

func init() { //nolint:gochecknoinits
	operateCmd.Flags().StringVarP(&action, "action", "a", "", "action (next|pause|play|play-pause|previous)")

	rootCmd.AddCommand(operateCmd)
}
