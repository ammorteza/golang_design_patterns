package main

import "fmt"

type Repair struct {
	next Representation
}

func NewRepair() *Repair{
	return &Repair{}
}

func (r *Repair)SetNext(next Representation){
	r.next = next
}

func (r *Repair)Execute(car *Car)  {
	if car.repairDone {
		fmt.Println("the car already was repaired!")
	}else{
		car.repairDone = true
		fmt.Println("the car is repaired.")
	}

	r.next.Execute(car)
}