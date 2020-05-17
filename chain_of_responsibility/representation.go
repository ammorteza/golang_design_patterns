package main

type Representation interface {
	Execute(car *Car)
	SetNext(representation Representation)
}