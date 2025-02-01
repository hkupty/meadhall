package meadhall

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"time"

	"github.com/hkupty/meadhall/api"
	"github.com/hkupty/meadhall/pkg/meadhall/config"
	"github.com/hkupty/meadhall/pkg/meadhall/wayland"
)

var (
	signals chan os.Signal
	done    chan bool
)

type App struct {
	socket     net.Listener
	appContext context.Context
	wayland    wayland.AppState
}

func Main(appContext context.Context) {
	waylandApp := wayland.NewApp()
	err := waylandApp.InitWayland()

	if err != nil {
		slog.Error("Could not connect to wayland, aborting", "error", err)
		os.Exit(1)
	}

	context.AfterFunc(appContext, waylandApp.Cleanup)

	sock, err := api.Socket(appContext, api.Daemon)

	if err != nil {
		slog.Error("Failed to start unix socket, thus unable to proceed", "error", err)
		os.Exit(1)
	}

	defer sock.Close()
	context.AfterFunc(appContext, func() { api.CleanupSocket(api.Daemon) })

	app := App{socket: sock, appContext: appContext, wayland: waylandApp}

	app.serve()
}

func (a *App) serve() {
	cfg, err := config.LoadConfig()

	if err != nil {
		slog.Warn("Failed to get config, proceeding with defaults")
	}

	// Init wayland handlers async
	go func() {
		for !a.wayland.Ready() {
		}
	}()

}

func registerIdleHandlers(cfg []config.IdleConfigItem, app *wayland.AppState) {
	var retries int = 10

	for !app.Ready() {
		time.Sleep(1 * time.Second)
		retries -= 1
		if retries == 0 {
			done <- true
		}
	}

	fmt.Println("Registering idle handlers")

	for _, idleConfig := range cfg {
		var idleHandler wayland.IdleEventHandler
		var resumedHandler wayland.IdleEventHandler

		switch idleConfig.Action.Type {
		case "output":
			idleHandler, resumedHandler = outputHandler(idleConfig, app)

		case "cmd":
			idleHandler, resumedHandler = cmdHandler(idleConfig)
		}

		app.RegisterNewIdleEventHandler(idleConfig.Timeout*1000, idleHandler, resumedHandler)
	}
}
