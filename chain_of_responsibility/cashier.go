package main

import "fmt"

type Cashier struct {
	next Representation
}

func NewCashier() *Cashier{
	return &Cashier{}
}

func (c *Cashier)SetNext(next Representation){
	c.next = next
}

func (c *Cashier)Execute(car *Car){
	if car.cashierDone {
		fmt.Println("the car already was cashiered!")
	}else{
		car.cashierDone = true
		fmt.Println("the car is cashiered.")
	}
	c.next.Execute(car)
}