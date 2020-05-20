package main

import "fmt"

func main(){
	server := NewNginx()
	fmt.Println(server.httpHandler("/users", "GET"))
	fmt.Println(server.httpHandler("/users", "GET"))
	fmt.Println(server.httpHandler("/users", "GET"))

	fmt.Println(server.httpHandler("/user/register", "POST"))
	fmt.Println(server.httpHandler("/user/register", "POST"))
}