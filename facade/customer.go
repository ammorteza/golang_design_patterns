package main

import (
	"errors"
	"fmt"
)

type Customers struct {
	users 			[]User
}

var users []User

func init(){
	users = append(users, NewUser("morteza amzajerdi"))
	users = append(users, NewUser("ali jahanpak"))
	users = append(users, NewUser("ali amzajerdi"))
	users = append(users, NewUser("abas amzajerdi"))
}

func NewCustomers() *Customers{
	fmt.Println("making a new customers unit")
	cm := &Customers{
		users: users,
	}

	fmt.Println("customers unit was made")
	return cm
}

func (this *Customers)IsValidUser(name string) error{
	fmt.Println("start validating user.")
	for _, user := range this.users{
		if user.Name == name{
			fmt.Println(name , " is a valid user")
			return nil
		}
	}

	fmt.Println(name, " is not a valid user")
	return errors.New("the user is not valid user")
}