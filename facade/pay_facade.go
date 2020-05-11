package main

import "fmt"

type PayFacade struct {

}

func NewPayFacade() *PayFacade{
	fmt.Println("making a new pay facade")
	pf := &PayFacade{}
	fmt.Println("pay facade was made")
	return pf
}

func (this *PayFacade)AddToList() error {
	fmt.Println("passing request to pay facade ...")
	return nil
}