// Meadhall daemon, runs indefinitely until receiving a signal to stop
// This package doesn't hold any logic and only deals with handling the binary-to-OS interface.
// Most notably, the signal handling and any parameter to config (eventual CLI flags or environment variables)
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/hkupty/meadhall/meadhall/pkg"
)

func main() {
	appContext, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pkg.Main(appContext)

	<-appContext.Done()
}
