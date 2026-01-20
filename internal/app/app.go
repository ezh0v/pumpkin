package app

type Instance struct {
	Version  string
	database Database
}

func New(version string, database Database) *Instance {
	return &Instance{
		Version:  version,
		database: database,
	}
}
