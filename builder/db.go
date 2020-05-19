package main

type db struct {
	iDatabase
}

func NewDb(d iDatabase) *db{
	return &db{
		d,
	}
}

func (d *db)setDatabase(dd iDatabase){
	d.iDatabase = dd
}

func (d *db)buildDatabase() Database{
	d.SetExpiration()
	d.SetConnectionString()
	return d.getDatabase()
}

