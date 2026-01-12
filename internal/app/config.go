package app

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Config struct {
	Version         string
	DatabaseConnect string
}

func (c *Config) validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Version, validation.NilOrNotEmpty.Error("unknown")),
		validation.Field(&c.DatabaseConnect, validation.NilOrNotEmpty.Error("empty value")),
	)
}
