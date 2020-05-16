package main

import (
	"fmt"
	"log"
)

func main(){
	circle, err := GetShape("circle", "first circle")
	if err != nil{
		log.Printf(err.Error())
	}
	square, err := GetShape("square", "first square")
	if err != nil{
		log.Printf(err.Error())
	}

	fmt.Println(circle.Name())
	fmt.Println(square.Name())
}