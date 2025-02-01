package api

import (
	"context"
	"log/slog"
	"net"
	"os"
)

func Socket(appContext context.Context, target Target) (net.Listener, error) {
	path, err := SocketPath(target)
	if err != nil {
		return nil, err
	}

	var listenConfig net.ListenConfig

	return listenConfig.Listen(appContext, "unix", path)
}

func CleanupSocket(target Target) {
	path, err := SocketPath(target)
	if err != nil {
		slog.Warn("Failed to acquire path for socket", "error", err)
	}

	err = os.Remove(path)
	slog.Warn("Failed to remove socket", "error", err)
}
