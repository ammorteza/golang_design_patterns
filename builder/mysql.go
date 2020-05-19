package main

type mysql struct {
	exp 					int
	connectionString 		string
}

func (m *mysql)SetExpiration(){
	m.exp = 10
}

func (m *mysql)SetConnectionString(){
	m.connectionString = "mysql connection string"
}

func (m *mysql)getDatabase() Database{
	return Database{
		exp: m.exp,
		connectionString: m.connectionString,
	}
}