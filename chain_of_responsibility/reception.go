package main

import "fmt"

type Reception struct {
	next 	Representation
}

func NewReception() *Reception{
	return &Reception{}
}

func (r *Reception)SetNext(next Representation){
	r.next = next
}

func (r *Reception)Execute(car *Car){
	if car.receptionDone {
		fmt.Println("the car already was received!")
	}else{
		car.receptionDone = true
		fmt.Println("the car is received.")
	}
	r.next.Execute(car)
}