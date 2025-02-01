package api

import (
	"context"
	"log/slog"
	"net"
)

type App struct {
	Listener net.Listener
	Context  context.Context
	Feature  Feature
}

func NewApp(appContext context.Context, feature Feature) (app *App, err error) {
	sock, err := FeatureListener(appContext, feature)

	if err != nil {
		slog.Error("Failed to start unix socket, thus unable to proceed", "error", err)
		return nil, err
	}

	context.AfterFunc(appContext, func() {
		sock.Close()
		CleanupSocket(feature)
	})

	app = &App{
		Listener: sock,
		Context:  appContext,
		Feature:  feature,
	}

	return
}
