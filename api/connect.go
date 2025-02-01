package api

import (
	"net"
)

func PrepareSend(message MessageType) (net.Conn, *Envelope, error) {
	envelope := Envelope{Type: message}
	socket, err := PrepareSendEnvelope(&envelope)

	if err != nil {
		return nil, nil, err
	}

	return socket, &envelope, nil
}

func PrepareSendEnvelope(envelope *Envelope) (net.Conn, error) {

	socketPath, err := SocketPath(Route(envelope.Type))
	if err != nil {
		return nil, err
	}

	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
