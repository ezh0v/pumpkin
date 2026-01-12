package server

type Option func(*options)

type options struct {
	address string
}

func optionsWithDefaults(opts *options) {
	if opts.address == "" {
		opts.address = ":8000"
	}
}

func WithAddress(val string) Option {
	return func(o *options) {
		o.address = val
	}
}
