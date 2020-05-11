package main

import "fmt"

type MarketFacade struct {
	customers 		*Customers
	payFacade 		*PayFacade
	storeroom 		*StoreRoom
}

func NewMarketFacade() *MarketFacade{
	fmt.Println("making a new market facade")
	mf := &MarketFacade{
		customers: NewCustomers(),
		payFacade: NewPayFacade(),
		storeroom: NewStoreRoom("storeroom 1"),
	}
	fmt.Println("market facade was made")
	return mf
}

func (this *MarketFacade)Order(uName string, pName string, count uint) {
	if err := this.customers.IsValidUser(uName); err != nil{
		fmt.Printf("error in order product: %s", err.Error())
		return
	}

	if err := this.storeroom.CheckExist(pName, count); err != nil{
		fmt.Printf("error in order product: %s", err.Error())
		return
	}

	if err := this.storeroom.ReserveProduct(pName, count); err != nil{
		fmt.Printf("error in order product: %s", err.Error())
		return
	}

	if err := this.payFacade.AddToList(); err != nil{
		fmt.Printf("error in order product: %s", err.Error())
		return
	}

	fmt.Println("your request registered successfully.")
}