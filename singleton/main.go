package main

import "fmt"

func main()  {
	for i := 0 ;i < 10; i++{
		go GetTypicalSingle()
		go GetSyncOnceInstance()
	}

	fmt.Scanln()
}