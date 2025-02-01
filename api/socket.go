package api

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"os"

	"github.com/adrg/xdg"
)

var (
	BadTarget = errors.New("Unknown Target")
)

func socketFileForFeature(target Feature) (path string, err error) {
	switch target {
	case Manager:
		path = "meadhall.sock"
	case Idle:
		path = "idle.sock"
	default:
		err = BadTarget
	}

	return
}

func socketPath(target Feature) (string, error) {
	path, err := socketFileForFeature(target)

	if err != nil {
		return "", err
	}

	return xdg.RuntimeFile("meadhall/" + path)
}

func FeatureListener(appContext context.Context, target Feature) (net.Listener, error) {
	path, err := socketPath(target)
	if err != nil {
		return nil, err
	}

	var listenConfig net.ListenConfig

	return listenConfig.Listen(appContext, "unix", path)
}

func CleanupSocket(target Feature) {
	path, err := socketPath(target)
	if err != nil {
		slog.Warn("Failed to acquire path for socket", "error", err)
	}

	err = os.Remove(path)
	slog.Warn("Failed to remove socket", "error", err)
}
