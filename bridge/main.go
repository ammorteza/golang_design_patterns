package main

import "fmt"

type Vehicle interface {
	Start()
	ChangeSupplier(Supplier)
}

type Supplier interface {
	Supply() 	string
}

type Petrol struct {}
type Gas struct {}

func (this Petrol)Supply() string{
	return "engine started and is working with petrol."
}

func (this Gas)Supply()  string{
	return "engine started and is working with gas."
}

func NewPetrol() Supplier{
	return Petrol{}
}

func NewGas() Supplier {
	return Gas{}
}

type Samand struct {
	PowerSupplier 		Supplier
	Name 				string
}

func NewSamand(supplier Supplier, name string) Vehicle{
	return &Samand{
		PowerSupplier: supplier,
		Name: name,
	}
}

func (this *Samand)Start(){
	fmt.Printf("the car %s %s \n" , this.Name, this.PowerSupplier.Supply())
}

func (this *Samand)ChangeSupplier(supplier Supplier) {
	this.PowerSupplier = supplier
}

func main(){
	vehicle := NewSamand(NewPetrol(), "samand ef7")
	vehicle.Start()
	vehicle.ChangeSupplier(NewGas())
	vehicle.Start()
}