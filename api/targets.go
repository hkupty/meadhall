package api

//go:generate msgp

type Target uint8
type Lifecycle uint8

const (
	Daemon Target = iota
	Idle
)

const (
	Started Lifecycle = iota
	Finished
)

// Identifies the right destination for the message type
func Route(msgType MessageType) Target {

	switch msgType {

	// Messages handled by the idle daemon
	case ListIdle:
		return Idle

	}

	// Daemon is the catch-all in case we have an unhandled message type
	return Daemon
}

func (t Target) Name() string {
	switch t {
	case Daemon:
		return "Daemon"
	case Idle:
		return "Idle"
	}

	return "unknown"
}
