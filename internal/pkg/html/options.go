package html

type Option func(*options)

type options struct {
	globalValues map[string]any
}

func WithGlobalValue(key string, val any) Option {
	return func(o *options) {
		o.globalValues[key] = val
	}
}
