package main

import "fmt"

type Release struct {
	next Representation
}

func NewRelease()  *Release{
	return &Release{}
}

func (r *Release)SetNext(next Representation){
	r.next = next
}

func (r *Release)Execute(car *Car){
	if car.releaseDone {
		fmt.Println("the car already was released!")
	}else{
		car.releaseDone = true
		fmt.Println("the car is released.")
	}
}