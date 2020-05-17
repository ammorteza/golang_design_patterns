package main

type Car struct {
	name 			string
	receptionDone 	bool
	repairDone 		bool
	cashierDone 	bool
	releaseDone		bool
}

func NewCar(name string) *Car{
	return &Car{
		name: name,
	}
}