package main

func main(){
	car := NewCar("peugeot 207")
	// handlers
	release := NewRelease()
	cashier := NewCashier()
	repair := NewRepair()
	reception := NewReception()

	cashier.SetNext(release)
	repair.SetNext(cashier)
	reception.SetNext(repair)
	reception.Execute(car)
}