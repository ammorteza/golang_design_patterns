package main

import "fmt"

type dbConnection struct {
	id 				int
	connection 		string
}

func NewDbConnection(id int) iPoolObject{
	return dbConnection{
		id: id,
		connection: fmt.Sprint("connection for object ", id),
	}
}

func (c dbConnection)GetId() int{
	return c.id
}

func (c dbConnection)GetConnection() string{
	return c.connection
}