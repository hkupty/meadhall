package api

import "github.com/tinylib/msgp/msgp"

//go:generate msgp
//msgp:tuple Envelope NotifyBarrelStatus

type MessageType uint8

const (
	ListIdle MessageType = iota
	Status
	BarrelStatus
)

type Envelope struct {
	Type MessageType
	Data msgp.Raw
}

// Idle handling

type IdleItem struct {
	Delay        uint32 `msg:"delay"`
	IdleAction   string `msg:"idle"`
	ResumeAction string `msg:"resume"`
}

type RequestListIdle struct{}
type ResponseListIdle struct {
	Items []IdleItem `msg:"items"`
}

// General messages
type RequestStatus struct{}
type ResponseStatus struct {
	Enabled []string `msg:"enabled"`
}

type NotifyBarrelStatus struct {
	Target    Target
	Lifecycle Lifecycle
}
