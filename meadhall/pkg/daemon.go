package pkg

import (
	"context"
	"log/slog"
	"os"

	"github.com/hkupty/meadhall/api"
	"github.com/hkupty/meadhall/api/config"
)

func Main(appContext context.Context) {
	app, err := api.NewApp(appContext, api.Manager)

	if err != nil {
		slog.Error("Failed to start unix socket, thus unable to proceed", "error", err)
		os.Exit(1)
	}

	serve(app)
}

func serve(app *api.App) {
	_, err := config.LoadConfig()

	if err != nil {
		slog.Warn("Failed to get config, proceeding with defaults")
	}
}
