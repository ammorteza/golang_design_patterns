package main

import "fmt"

type TV struct {}

func (this *TV)Plug220Volt() string{
	return fmt.Sprint("turned on TV with 220 voltage.")
}