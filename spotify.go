package spotify

import (
	"github.com/godbus/dbus"
)

type CLI struct {
	conn *dbus.Conn
}

func New() (*CLI, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}
	return &CLI{conn: conn}, nil
}
