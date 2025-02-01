package api

// Each binary should be responsible for a single feature.
// Therefore, the feature is the global locator for message exchange.
// Additionally, the manager (meadhall daemon) should know when and how to route messages across
// features
type Feature uint8

const (
	// The Main process, a daemon that runs indefinitely and listens to CLI messages as well as
	// notifications from each feature.
	// Features can exchange messages indirectly through Manager as well.
	// The Manager itself isn't a wayland client and is only responsible for ensuring the clients
	// are healthy and active
	Manager Feature = iota

	// Idle manager: Responsible for triggering actions after some idle time.
	// Additionally, some cleanup action can be executed when activity is restored.
	// Uses `ext-idle-notify-v1` from wayland-protocols
	Idle
)
