package daemon

import "time"

// Daemon holds controls
type Daemon interface {
	Start() <-chan error
	Stop()
}

type daemon struct {
	action   Action
	errCh    <-chan error
	interval time.Duration
	ticker   *time.Ticker
}

// Action is a user-defined function
type Action func() error

// Start starts the daemon and gives the user a read
// only error channel for error handling
func (d *daemon) Start() <-chan error {
	d.ticker = time.NewTicker(s.interval)
	d.errCh = make(chan error)

	go func() {
		for _ := range d.ticker.C {
			err := d.action()
			if err != nil {
				d.errCh <- err
			}
		}
	}()

	return errCh
}

// Stop stops the daemon by closing the ticker channel
// and closing the error channel
func (d *daemon) Stop() {
	d.ticker.Stop()
	close(d.errCh)
}
