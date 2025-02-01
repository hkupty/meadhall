package main

import (
	"fmt"
	"net"

	"github.com/hkupty/meadhall/api"
	"github.com/tinylib/msgp/msgp"
)

func main() {
	socketPath, err := api.SocketPath(api.Route(api.Status))
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")
	w := msgp.NewWriter(conn)

	statusRequest := api.RequestStatus{}
	data, err := statusRequest.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}
	envelope := api.Envelope{Type: api.Status, Data: data}

	err = envelope.EncodeMsg(w)
	if err != nil {
		panic(err)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}
	fmt.Println("sent request")

	err = envelope.DecodeMsg(msgp.NewReader(conn))
	if err != nil {
		panic(err)
	}
	fmt.Println("got response")
	var response api.ResponseStatus

	_, err = response.UnmarshalMsg(envelope.Data)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Enabled)
}
