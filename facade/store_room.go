package main

import (
	"errors"
	"fmt"
)

type StoreRoom struct {
	Name 			string
	Products 		[]Product
}

var products []Product

func init(){
	products = append(products, NewProduct("tv-lg", 10))
	products = append(products, NewProduct("laptop-dell", 100))
	products = append(products, NewProduct("radio-23gth", 15))
	products = append(products, NewProduct("camera-sony", 5))
	products = append(products, NewProduct("dish-dd", 190))
}

func NewStoreRoom(name string) *StoreRoom{
	fmt.Println("making a storeroom")
	storeroom := &StoreRoom{
		Name: name,
		Products: products,
	}
	fmt.Println("storeroom was made")
	return storeroom
}

func (this *StoreRoom)CheckExist(pName string, count uint) error{
	fmt.Println("start checking exist " , pName, " in storeroom")
	for _, product := range this.Products{
		if product.Name == pName && product.Count >= count{
			fmt.Println(pName, " exist in storeroom")
			return nil
		}
	}

	return errors.New(pName + " is not exist in storeroom")
}

func (this *StoreRoom)ReserveProduct(pName string, count uint) error{
	fmt.Println("start reserving product")
	for _, product := range this.Products{
		if product.Name == pName{
			product.Count -= count
			return nil
		}
	}

	return errors.New("couldn't find product in storeroom")
}