package wayland

import (
	"errors"
	"fmt"
	"maps"
	"slices"

	"github.com/hkupty/meadhall/pkg/meadhall/wayland/gen"
	"github.com/rajveermalviya/go-wayland/wayland/client"
)

func (app *AppState) InitWayland() error {
	display, err := client.Connect("")

	if err != nil {
		// TODO proper log
		fmt.Printf("Failed to connect to wayland server/compositor: %v \n", err)
		return err
	}

	app.Display = display

	registry, err := display.GetRegistry()

	if err != nil {
		fmt.Printf("Failed to connect to get registry: %v \n", err)
		return err
	}

	app.registry = registry

	app.registry.SetGlobalHandler(app.HandleRegistryEvents)

	return nil
}

func (app *AppState) StartEventLoop() error {
	for {
		if err := app.Context().Dispatch(); err != nil {
			return err
		}
	}
}

/*
This is where all the supported protocols are registered so the wayland compositor knows what kinds of messages
we can handle.
*/
func (app *AppState) HandleRegistryEvents(e client.RegistryGlobalEvent) {
	fmt.Printf("Got event %v\n", e)
	switch e.Interface {
	case "wl_seat":
		seat := client.NewSeat(app.Context())
		err := app.registry.Bind(e.Name, e.Interface, e.Version, seat)
		if err != nil {
			panic(err)
		}
		app.seat = seat
		// TODO handle multi-seat?

	case "ext_idle_notifier_v1":
		notifier := gen.NewExtIdleNotifierV1(app.Context())
		err := app.registry.Bind(e.Name, e.Interface, e.Version, notifier)
		if err != nil {
			panic(err)
		}
		app.idleNotifier = notifier

	case "wl_output":
		output := client.NewOutput(app.Context())
		err := app.registry.Bind(e.Name, e.Interface, e.Version, output)
		if err != nil {
			panic(err)
		}
		output.SetNameHandler(app.getOutputNameHandler(output))

	case "zwlr_output_power_manager_v1":
		powerManager := gen.NewZwlrOutputPowerManagerV1(app.Context())
		err := app.registry.Bind(e.Name, e.Interface, e.Version, powerManager)
		if err != nil {
			panic(err)
		}
		app.outputPowerManager = powerManager

	}

	// TODO zwlr_layer_shell_v1 support for overlays/status bars?
	// TODO zwlr_foreign_toplevel_manager_v1 for status bar
	// TODO ext_session_lock_manager_v1 for locking the session
}

func (app *AppState) getOutputNameHandler(output *client.Output) func(client.OutputNameEvent) {
	return func(event client.OutputNameEvent) {
		output := OutputWrapper{output, nil, false}
		app.outputs[event.Name] = &output
		fmt.Printf("Got a new output name: %v\n", event.Name)
	}
}

func (app *AppState) RegisterNewIdleEventHandler(timeout uint32, idleHandler IdleEventHandler, resumedHandler IdleEventHandler) error {

	if app.seat == nil {
		return errors.New("seat: Wayland seat not registered to the app")
	}

	notificationHandler, err := app.idleNotifier.GetIdleNotification(timeout, app.seat)

	if err != nil {
		return err
	}

	if idleHandler != nil {
		notificationHandler.SetIdledHandler(func(e gen.ExtIdleNotificationV1IdledEvent) {
			err := idleHandler()
			if err != nil {
				fmt.Printf("Due to errors, this idle handler is being removed: %v", err)
				notificationHandler.Destroy()
				if resumedHandler != nil {
					err := resumedHandler()
					if err != nil {
						fmt.Printf("Attempted to restore after removing notifier, but an error occurred: %v", err)
					}
				}
			}
		})
	}

	if resumedHandler != nil {
		notificationHandler.SetResumedHandler(func(e gen.ExtIdleNotificationV1ResumedEvent) {
			err := resumedHandler()
			if err != nil {
				fmt.Printf("Due to errors, this idle handler is being removed: %v", err)
				notificationHandler.Destroy()
			}
		})
	}

	// TODO hold notification handler

	fmt.Printf("Registered notification handler (%d ms)", timeout)

	return nil
}

func (app *AppState) SetOutputPower(outputName string, powered bool) error {
	output := app.outputs[outputName]

	if output.Output == nil {
		return fmt.Errorf("output: No such output with name %s", outputName)
	}

	if output.powerManager == nil {
		if app.outputPowerManager == nil {
			return errors.New("output: Cannot change power state since zwlr_output_power_manager_v1 is not registered")
		}

		outputPower, err := app.outputPowerManager.GetOutputPower(output.Output)
		if err != nil {
			return err
		}

		output.powerManager = outputPower
	}

	output.powerManager.SetModeHandler(output.SetCurrentPowerState)

	if powered {
		output.PowerOn()
	} else {
		output.PowerOff()
	}

	return nil
}

func (app *AppState) GetRegisteredOutputs() []string {
	keys := maps.Keys(app.outputs)
	return slices.Collect(keys)
}

func (app *AppState) Ready() bool {
	fmt.Printf("Ready status: display -> %t, registry -> %t, seat -> %t, idle -> %t, outputPowerManager -> %t\n",
		app.Display != nil,
		app.registry != nil,
		app.seat != nil,
		app.idleNotifier != nil,
		app.outputPowerManager != nil)
	return app.Display != nil && app.registry != nil && app.seat != nil && app.idleNotifier != nil && app.outputPowerManager != nil
}
