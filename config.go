package daemon

import "time"

// Config holds configuration for the daemon
type Config struct {
	Action   Action
	Interval time.Duration
}

// New returns a new Daemon
func New(config Config) Daemon {
	return &daemon{
		interval: config.Interval,
		action:   config.Action,
	}
}
