package app

type Instance struct {
	Version  string
	database Database

	admin *Admin
}

func New(version string, database Database) *Instance {
	return &Instance{
		Version:  version,
		database: database,
	}
}

func (i *Instance) Admin() *Admin {
	if i.admin == nil {
		i.admin = &Admin{Instance: i}
	}

	return i.admin
}
