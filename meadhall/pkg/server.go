package meadhall

import (
	"log"
	"maps"
	"net"
	"os"
	"os/signal"
	"slices"
	"syscall"

	"github.com/hkupty/meadhall/api"
	"github.com/tinylib/msgp/msgp"
)

var connected map[string]bool = make(map[string]bool)

func serve() error {
	path, err := api.SocketPath(api.Daemon)
	if err != nil {
		return err
	}

	println(path)

	socket, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}

	interrupChan := make(chan os.Signal, 1)
	signal.Notify(interrupChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupChan
		os.Remove(path)
		os.Exit(1)
	}()

	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()
			var envelope api.Envelope
			writer := msgp.NewWriter(conn)
			err := envelope.DecodeMsg(msgp.NewReader(conn))
			if err != nil {
				log.Fatal(err)
				return
			}
			log.Println("Got message")
			switch envelope.Type {
			case api.Status:
				log.Println("Status request")
				var request api.RequestStatus
				_, err = request.UnmarshalMsg([]byte(envelope.Data))
				if err != nil {
					log.Fatal(err)
					return
				}
				response := api.ResponseStatus{
					Enabled: slices.Collect(maps.Keys(connected)),
				}
				wrapped, err := response.MarshalMsg(nil)
				if err != nil {
					log.Fatal(err)
					return
				}
				envelope.Data = wrapped
				err = envelope.EncodeMsg(writer)

				if err != nil {
					log.Fatal(err)
					return
				}

				writer.Flush()

			case api.BarrelStatus:
				log.Println("Barrel notification incoming")
				var notification api.NotifyBarrelStatus
				_, err = notification.UnmarshalMsg([]byte(envelope.Data))
				if err != nil {
					log.Fatal(err)
					return
				}

				target := notification.Target.Name()

				switch notification.Lifecycle {
				case api.Started:
					connected[target] = true
				case api.Finished:
					delete(connected, target)
				}
			}

		}(conn)

	}

}
