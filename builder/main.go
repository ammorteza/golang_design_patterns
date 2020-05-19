package main

import "fmt"

func main(){
	mysql := NewDatabase("mysql")

	db := NewDb(mysql)
	dbConn := db.buildDatabase()
	fmt.Println(dbConn.connectionString, " , expiration time ", dbConn.exp)

	postgresql := NewDatabase("postgresql")
	db.setDatabase(postgresql)
	dbConn = db.buildDatabase()
	fmt.Println(dbConn.connectionString, " , expiration time ", dbConn.exp)
}