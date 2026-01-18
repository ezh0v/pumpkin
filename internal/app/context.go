package app

type Context struct {
	Version  string
	database Database
}

func NewContext(version string, database Database) *Context {
	return &Context{
		Version:  version,
		database: database,
	}
}
