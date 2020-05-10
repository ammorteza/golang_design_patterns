package main

import "fmt"

type Device interface {
	Plug220Volt()	string
}

type Person struct {
	Name 		string
}

func  NewPerson(name string) *Person{
	return &Person{
		name,
	}
}

func (this *Person)TurnOnDevice(device Device)  {
	fmt.Println(this.Name + " " + device.Plug220Volt())
}