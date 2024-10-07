package meadhall

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/hkupty/meadhall/pkg/meadhall/config"
	"github.com/hkupty/meadhall/pkg/meadhall/wayland"
)

// Returns a pair of [wayland.IdleEventHandler] for when idling and resuming that turns on and off the
// outputs respectively. If none provided through [config.IdleAction.Args], will default to all outputs.
// Note that [config.IdleAction.ResumeArgs] are ignored for this handler
func outputHandler(idleConfig config.IdleConfigItem, app *wayland.AppState) (wayland.IdleEventHandler, wayland.IdleEventHandler) {
	var idleHandler wayland.IdleEventHandler
	var resumedHandler wayland.IdleEventHandler

	targets := idleConfig.Action.OnIdle

	if len(targets) == 0 {
		targets = app.GetRegisteredOutputs()
	}

	if len(idleConfig.Action.OnResume) > 0 {
		fmt.Printf("warn: ResumeArgs are not used by the \"output\" handler and will be ignored")
	}

	idleHandler = func() error {
		var errs error
		for _, output := range targets {
			fmt.Printf("Setting power off for output %s\n", output)
			err := app.SetOutputPower(output, false)
			if err != nil {
				errs = errors.Join(errs, err)
			}
		}

		return errs
	}

	resumedHandler = func() error {
		var errs error
		for _, output := range targets {
			fmt.Printf("Setting power on for output %s\n", output)
			err := app.SetOutputPower(output, true)
			if err != nil {
				errs = errors.Join(errs, err)
			}
		}

		return errs
	}

	return idleHandler, resumedHandler
}

// Returns a pair of [wayland.IdleEventHandler] for when idling and resuming that executes commands upon those events are received
func cmdHandler(idleConfig config.IdleConfigItem) (wayland.IdleEventHandler, wayland.IdleEventHandler) {
	var idleHandler wayland.IdleEventHandler
	var resumedHandler wayland.IdleEventHandler
	args := idleConfig.Action.OnIdle
	resumeArgs := idleConfig.Action.OnResume

	if len(args) > 0 {
		idleHandler = func() error {
			fmt.Println("Idle Triggered")
			return runCmd(args)
		}
	}

	if len(resumeArgs) > 0 {
		resumedHandler = func() error {
			fmt.Println("Resume triggered")
			return runCmd(resumeArgs)
		}
	}

	return idleHandler, resumedHandler
}

func runCmd(cmdWithArgs []string) error {
	cmd := exec.Command(cmdWithArgs[0], cmdWithArgs[1:]...)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		panic(err)
	}

	go func() {
		io.Copy(os.Stdout, stdout)
	}()

	err = cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}
