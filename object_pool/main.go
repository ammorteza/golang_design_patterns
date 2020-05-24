package main

import (
	"fmt"
	"log"
)

func main(){
	objects := make([]iPoolObject, 0)

	for i := 0; i < 3; i++{
		objects = append(objects, NewDbConnection(i))
	}

	pool, err := initPool(objects)
	if err != nil{
		log.Fatal(err)
	}

	obj, err := pool.Loan()
	if err != nil{
		log.Fatal(err)
	}

	obj1, err := pool.Loan()
	if err != nil{
		log.Fatal(err)
	}


	fmt.Println(obj.GetConnection())
	fmt.Println(obj1.GetConnection())

	if err := pool.receive(obj); err != nil{
		log.Fatal(err)
	}
	if err := pool.receive(obj1); err != nil{
		log.Fatal(err)
	}
}