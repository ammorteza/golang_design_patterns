package main

import (
	"fmt"
	"sync"
)

var (
	typicalSingleInstace 	*single
	lock 					sync.Mutex
)

func GetTypicalSingle() *single{
	if typicalSingleInstace == nil{
		lock.Lock()
		defer lock.Unlock()
		if typicalSingleInstace == nil{
			fmt.Println("an instance of single struct was made.")
			typicalSingleInstace = &single{}
		}else {
			fmt.Println("an instance of single struct already madden.")
		}
	}else {
		fmt.Println("an instance of single struct already madden.")
	}

	return typicalSingleInstace
}