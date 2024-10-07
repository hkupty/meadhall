package meadhall

import (
	"fmt"
	"sync"
	"time"

	"github.com/hkupty/meadhall/pkg/meadhall/config"
	"github.com/hkupty/meadhall/pkg/meadhall/wayland"
)

func Main() {
	var waitGroup sync.WaitGroup
	cfg := config.LoadConfig()

	fmt.Println(cfg)
	fmt.Println(cfg.Idle)

	waitGroup.Add(1)
	app := connectWaylandClient()

	go func() {
		defer waitGroup.Done()
		for {
			if err := app.StartEventLoop(); err != nil {
				fmt.Printf("Got an error, finishing: %v", err)
				return
			}
		}
	}()

	registerIdleHandlers(cfg.Idle, app)

	waitGroup.Wait()
}

func connectWaylandClient() *wayland.AppState {
	waylandApp := wayland.NewApp()
	err := waylandApp.InitWayland()
	if err != nil {
		panic(err)
	}

	return &waylandApp

}

func registerIdleHandlers(cfg []config.IdleConfigItem, app *wayland.AppState) {

	for !app.Ready() {
		time.Sleep(1 * time.Second)
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
