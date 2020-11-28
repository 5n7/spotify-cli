package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version bool

func runRoot(cmd *cobra.Command, args []string) error {
	if version {
		return runVersion(cmd, args)
	}
	return fmt.Errorf("spotify-cli requires a sub-comand")
}

var rootCmd = &cobra.Command{
	Use:   "spotify-cli",
	Short: "spotify-cli is a DBus based Spotify client",
	Long:  "spotify-cli is a DBus based Spotify client.",
	RunE:  runRoot,
}

func init() { //nolint:gochecknoinits
	rootCmd.Flags().BoolVarP(&version, "version", "V", false, "show version")
}

func Execute() {
	rootCmd.Execute() //nolint:errcheck
}
