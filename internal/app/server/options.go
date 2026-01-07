package server

type Option func(*options)

type options struct {
	viewsPatterns string
	address       string
}

func optionsWithDefaults(opts *options) {
	if opts.address == "" {
		opts.address = ":8000"
	}

	if opts.viewsPatterns == "" {
		opts.viewsPatterns = "views/**/*.html"
	}
}

func WithViewsPatterns(val string) Option {
	return func(o *options) {
		o.viewsPatterns = val
	}
}

func WithAddress(val string) Option {
	return func(o *options) {
		o.address = val
	}
}
