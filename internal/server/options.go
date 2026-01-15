package server

import "time"

type Option func(*options)

type options struct {
	address         string
	shutdownTimeout time.Duration
}

func (opts *options) withDefaults() {
	if opts.address == "" {
		opts.address = ":8000"
	}

	if opts.shutdownTimeout == 0 {
		opts.shutdownTimeout = 5 * time.Second
	}
}

func WithAddress(val string) Option {
	return func(o *options) {
		o.address = val
	}
}

func WithShutdownTimeout(val time.Duration) Option {
	return func(o *options) {
		o.shutdownTimeout = val
	}
}
