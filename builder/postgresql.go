package main

type postgresql struct {
	exp 					int
	connectionString		string
}

func (p *postgresql)SetExpiration(){
	p.exp = 15
}

func (p *postgresql)SetConnectionString(){
	p.connectionString = "postgresql connection string"
}

func (p *postgresql)getDatabase() Database{
	return Database{
		exp: p.exp,
		connectionString: p.connectionString,
	}
}