package main

type iDatabase interface {
	SetExpiration()
	SetConnectionString()
	getDatabase() Database
}

func NewDatabase(t string) iDatabase{
	switch t {
	case "mysql":
		return &mysql{}
	case "postgresql":
		return &postgresql{}
	default:
		return nil
	}
}