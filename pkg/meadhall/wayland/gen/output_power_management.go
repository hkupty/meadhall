// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : protocols/wlr-output-power-management-unstable-v1.xml
//
// wlr_output_power_management_unstable_v1 Protocol Copyright:
//
// Copyright © 2019 Purism SPC
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice (including the next
// paragraph) shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package gen

import "github.com/rajveermalviya/go-wayland/wayland/client"

// ZwlrOutputPowerManagerV1 : manager to create per-output power management
//
// This interface is a manager that allows creating per-output power
// management mode controls.
type ZwlrOutputPowerManagerV1 struct {
	client.BaseProxy
}

// NewZwlrOutputPowerManagerV1 : manager to create per-output power management
//
// This interface is a manager that allows creating per-output power
// management mode controls.
func NewZwlrOutputPowerManagerV1(ctx *client.Context) *ZwlrOutputPowerManagerV1 {
	zwlrOutputPowerManagerV1 := &ZwlrOutputPowerManagerV1{}
	ctx.Register(zwlrOutputPowerManagerV1)
	return zwlrOutputPowerManagerV1
}

// GetOutputPower : get a power management for an output
//
// Create an output power management mode control that can be used to
// adjust the power management mode for a given output.
func (i *ZwlrOutputPowerManagerV1) GetOutputPower(output *client.Output) (*ZwlrOutputPowerV1, error) {
	id := NewZwlrOutputPowerV1(i.Context())
	const opcode = 0
	const _reqBufLen = 8 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], id.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], output.ID())
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}

// Destroy : destroy the manager
//
// All objects created by the manager will still remain valid, until their
// appropriate destroy request has been called.
func (i *ZwlrOutputPowerManagerV1) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 1
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// ZwlrOutputPowerV1 : adjust power management mode for an output
//
// This object offers requests to set the power management mode of
// an output.
type ZwlrOutputPowerV1 struct {
	client.BaseProxy
	modeHandler   ZwlrOutputPowerV1ModeHandlerFunc
	failedHandler ZwlrOutputPowerV1FailedHandlerFunc
}

// NewZwlrOutputPowerV1 : adjust power management mode for an output
//
// This object offers requests to set the power management mode of
// an output.
func NewZwlrOutputPowerV1(ctx *client.Context) *ZwlrOutputPowerV1 {
	zwlrOutputPowerV1 := &ZwlrOutputPowerV1{}
	ctx.Register(zwlrOutputPowerV1)
	return zwlrOutputPowerV1
}

// SetMode : Set an outputs power save mode
//
// Set an output's power save mode to the given mode. The mode change
// is effective immediately. If the output does not support the given
// mode a failed event is sent.
//
//	mode: the power save mode to set
func (i *ZwlrOutputPowerV1) SetMode(mode uint32) error {
	const opcode = 0
	const _reqBufLen = 8 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(mode))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// Destroy : destroy this power management
//
// Destroys the output power management mode control object.
func (i *ZwlrOutputPowerV1) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 1
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

type ZwlrOutputPowerV1Mode uint32

// ZwlrOutputPowerV1Mode :
const (
	// ZwlrOutputPowerV1ModeOff : Output is turned off.
	ZwlrOutputPowerV1ModeOff ZwlrOutputPowerV1Mode = 0
	// ZwlrOutputPowerV1ModeOn : Output is turned on, no power saving
	ZwlrOutputPowerV1ModeOn ZwlrOutputPowerV1Mode = 1
)

func (e ZwlrOutputPowerV1Mode) Name() string {
	switch e {
	case ZwlrOutputPowerV1ModeOff:
		return "off"
	case ZwlrOutputPowerV1ModeOn:
		return "on"
	default:
		return ""
	}
}

func (e ZwlrOutputPowerV1Mode) Value() string {
	switch e {
	case ZwlrOutputPowerV1ModeOff:
		return "0"
	case ZwlrOutputPowerV1ModeOn:
		return "1"
	default:
		return ""
	}
}

func (e ZwlrOutputPowerV1Mode) String() string {
	return e.Name() + "=" + e.Value()
}

type ZwlrOutputPowerV1Error uint32

// ZwlrOutputPowerV1Error :
const (
	// ZwlrOutputPowerV1ErrorInvalidMode : nonexistent power save mode
	ZwlrOutputPowerV1ErrorInvalidMode ZwlrOutputPowerV1Error = 1
)

func (e ZwlrOutputPowerV1Error) Name() string {
	switch e {
	case ZwlrOutputPowerV1ErrorInvalidMode:
		return "invalid_mode"
	default:
		return ""
	}
}

func (e ZwlrOutputPowerV1Error) Value() string {
	switch e {
	case ZwlrOutputPowerV1ErrorInvalidMode:
		return "1"
	default:
		return ""
	}
}

func (e ZwlrOutputPowerV1Error) String() string {
	return e.Name() + "=" + e.Value()
}

// ZwlrOutputPowerV1ModeEvent : Report a power management mode change
//
// Report the power management mode change of an output.
//
// The mode event is sent after an output changed its power
// management mode. The reason can be a client using set_mode or the
// compositor deciding to change an output's mode.
// This event is also sent immediately when the object is created
// so the client is informed about the current power management mode.
type ZwlrOutputPowerV1ModeEvent struct {
	Mode uint32
}
type ZwlrOutputPowerV1ModeHandlerFunc func(ZwlrOutputPowerV1ModeEvent)

// SetModeHandler : sets handler for ZwlrOutputPowerV1ModeEvent
func (i *ZwlrOutputPowerV1) SetModeHandler(f ZwlrOutputPowerV1ModeHandlerFunc) {
	i.modeHandler = f
}

// ZwlrOutputPowerV1FailedEvent : object no longer valid
//
// This event indicates that the output power management mode control
// is no longer valid. This can happen for a number of reasons,
// including:
// - The output doesn't support power management
// - Another client already has exclusive power management mode control
// for this output
// - The output disappeared
//
// Upon receiving this event, the client should destroy this object.
type ZwlrOutputPowerV1FailedEvent struct{}
type ZwlrOutputPowerV1FailedHandlerFunc func(ZwlrOutputPowerV1FailedEvent)

// SetFailedHandler : sets handler for ZwlrOutputPowerV1FailedEvent
func (i *ZwlrOutputPowerV1) SetFailedHandler(f ZwlrOutputPowerV1FailedHandlerFunc) {
	i.failedHandler = f
}

func (i *ZwlrOutputPowerV1) Dispatch(opcode uint32, fd int, data []byte) {
	switch opcode {
	case 0:
		if i.modeHandler == nil {
			return
		}
		var e ZwlrOutputPowerV1ModeEvent
		l := 0
		e.Mode = client.Uint32(data[l : l+4])
		l += 4

		i.modeHandler(e)
	case 1:
		if i.failedHandler == nil {
			return
		}
		var e ZwlrOutputPowerV1FailedEvent

		i.failedHandler(e)
	}
}
