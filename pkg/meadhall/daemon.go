package meadhall

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hkupty/meadhall/pkg/meadhall/config"
	"github.com/hkupty/meadhall/pkg/meadhall/wayland"
)

var (
	signals chan os.Signal
	done    chan bool
)

func Main() {
	signals = make(chan os.Signal, 1)
	done = make(chan bool, 1)
	cfg := config.LoadConfig()

	fmt.Println(cfg)
	fmt.Println(cfg.Idle)

	app := connectWaylandClient()
	go func() {
		for {
			if err := app.StartEventLoop(); err != nil {
				fmt.Printf("Got an error, finishing: %v", err)
				done <- true
				return
			}
		}
	}()

	registerIdleHandlers(cfg.Idle, app)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go cleanup(app)
	<-done
}

func connectWaylandClient() *wayland.AppState {
	waylandApp := wayland.NewApp()
	err := waylandApp.InitWayland()
	if err != nil {
		fmt.Println(err)
		done <- true
	}

	return &waylandApp

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

func cleanup(app *wayland.AppState) {
	signal := <-signals
	fmt.Printf("\nGot %v signal\n", signal)
	if app != nil {
		app.Cleanup()
	}
	done <- true
}
