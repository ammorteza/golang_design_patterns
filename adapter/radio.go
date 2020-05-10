package main

import "fmt"

type Radio struct {}

func (this *Radio)Plug12Volt() string{
	return fmt.Sprint("turned on Radio with 12 voltage.")
}