package pkg

import (
	"fmt"
	"time"

	"github.com/hkupty/meadhall/api/config"
	"github.com/hkupty/meadhall/api/wayland"
)

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
