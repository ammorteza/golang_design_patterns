package main

import (
	"fmt"
)

var (
	InitSingleInstance 		*single
)

func init() {
	if InitSingleInstance == nil{
		fmt.Println("initSingleInstance was made.")
		InitSingleInstance = &single{}
	}else{
		fmt.Println("initSingleInstance already was made.")
	}
}