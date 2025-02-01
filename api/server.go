package api

import (
	"errors"

	"github.com/adrg/xdg"
)

func targetToFile(target Target) (string, error) {
	switch target {
	case Daemon:
		return "daemon.sock", nil
	case Idle:
		return "idle.sock", nil
	}

	return "", errors.New("Unknown target")
}

func SocketPath(target Target) (string, error) {
	filepath, err := targetToFile(target)
	if err != nil {
		return "", nil
	}

	return xdg.RuntimeFile("meadhall/" + filepath)
}
