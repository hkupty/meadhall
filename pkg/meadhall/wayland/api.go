package wayland

import (
	"github.com/rajveermalviya/go-wayland/wayland/client"

	"github.com/hkupty/meadhall/pkg/meadhall/wayland/gen"
)

// Wrapper over a [client.Display] object
type AppState struct {
	appID string

	*client.Display
	registry *client.Registry

	seat *client.Seat

	// Enables setting up Idle/Resume hooks
	idleNotifier *gen.ExtIdleNotifierV1

	outputs map[string]*OutputWrapper

	outputPowerManager *gen.ZwlrOutputPowerManagerV1

	// TODO Implement support for a session locker
	// sessionLockManager *gen.ExtSessionLockManagerV1
}

// Wraps the Output object to allow for a better control over it, i.e. by controling its power directly
type OutputWrapper struct {
	*client.Output
	powerManager *gen.ZwlrOutputPowerV1
	powered      bool
}

type IdleEvent struct {
	timeout int
	action  string // TODO change to something better
}

func (o *OutputWrapper) PowerOn() {
	o.powerManager.SetMode(uint32(gen.ZwlrOutputPowerV1ModeOn))
}

func (o *OutputWrapper) PowerOff() {
	o.powerManager.SetMode(uint32(gen.ZwlrOutputPowerV1ModeOff))
}

func (o *OutputWrapper) SetCurrentPowerState(e gen.ZwlrOutputPowerV1ModeEvent) {
	o.powered = e.Mode == uint32(gen.ZwlrOutputPowerV1ModeOn)
}

func NewApp() AppState {
	newApp := AppState{
		appID:   "meadhall",
		outputs: make(map[string]*OutputWrapper),
	}

	return newApp
}

type IdleEventHandler func() error
