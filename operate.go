package spotify

import (
	"fmt"

	"github.com/dawidd6/go-spotify-dbus"
)

type Action int

const (
	ActionNext Action = iota
	ActionPause
	ActionPlay
	ActionPlayPause
	ActionPrevious
)

func (c *CLI) Operate(action Action) error {
	switch action {
	case ActionNext:
		return spotify.SendNext(c.conn)
	case ActionPause:
		return spotify.SendPause(c.conn)
	case ActionPlay:
		return spotify.SendPlay(c.conn)
	case ActionPlayPause:
		return spotify.SendPlayPause(c.conn)
	case ActionPrevious:
		return spotify.SendPrevious(c.conn)
	default:
		return fmt.Errorf("the action is not supported")
	}
}
