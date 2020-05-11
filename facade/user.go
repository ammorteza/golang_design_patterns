package main

import (
	"math/rand"
)

type User struct {
	ID 			uint
	Name 		string
}

func NewUser(name string) User{
	return User{
		ID: uint(rand.Intn(1000)),
		Name: name,
	}
}