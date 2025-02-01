package idle

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hkupty/meadhall/api"
	"github.com/hkupty/meadhall/pkg/meadhall/config"
	"github.com/tinylib/msgp/msgp"
)

var (
	signals              chan os.Signal
	done                 chan bool
	notificationEnvelope api.Envelope = api.Envelope{Type: api.BarrelStatus}
	startedNotification  msgp.Raw
	finishedNotification msgp.Raw
)

//	func init() {
//		base := api.NotifyBarrelStatus{
//			Target:    api.Idle,
//			Lifecycle: api.Started,
//		}
//		var err error
//
//		startedNotification, err = base.MarshalMsg(nil)
//		if err != nil {
//			panic(err)
//		}
//
//		base.Lifecycle = api.Finished
//
//		finishedNotification, err = base.MarshalMsg(nil)
//		if err != nil {
//			panic(err)
//		}
//	}
func Main() {
	signals = make(chan os.Signal, 1)
	done = make(chan bool, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	cfg := config.LoadConfig()
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

	go func() {
		<-signals
		notifyShutdown()
		done <- true
	}()

	notifyStartup()

	failed := <-done
	if failed {
		os.Exit(1)
	}
}

func notifyShutdown() {
	conn, err := api.PrepareSendEnvelope(&notificationEnvelope)

	if err != nil {
		panic(err)
	}

	writer := msgp.NewWriter(conn)
	notificationEnvelope.Data = finishedNotification
	err = notificationEnvelope.EncodeMsg(writer)

	if err != nil {
		panic(err)
	}

	if err = writer.Flush(); err != nil {
		panic(err)
	}
}

func notifyStartup() {
	conn, err := api.PrepareSendEnvelope(&notificationEnvelope)

	if err != nil {
		panic(err)
	}

	writer := msgp.NewWriter(conn)
	notificationEnvelope.Data = startedNotification
	err = notificationEnvelope.EncodeMsg(writer)

	if err != nil {
		panic(err)
	}

	if err = writer.Flush(); err != nil {
		panic(err)
	}
}
