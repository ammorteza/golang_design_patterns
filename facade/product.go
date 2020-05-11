package main

import (
	"math/rand"
)

type Product struct {
	ID 			uint
	Name 		string
	Count 		uint
}

func NewProduct(name string, count uint) Product{
	return Product{
		ID: uint(rand.Intn(1000)),
		Name: name,
		Count: count,
	}
}