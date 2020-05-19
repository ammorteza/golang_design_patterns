package main

import (
	"fmt"
	"sync"
)

var (
	syncOnceSingleInstance 			*single
	once 							sync.Once
)

func GetSyncOnceInstance() *single{
	if syncOnceSingleInstance == nil{
		once.Do(func() {
			fmt.Println("an instance of syncOnceSingleInstance was made.")
			syncOnceSingleInstance = &single{}
		})
	}else{
		fmt.Println("an instance of syncOneSingleInstance already madden")
	}
	return syncOnceSingleInstance
}