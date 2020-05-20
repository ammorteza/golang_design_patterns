package main

type Server interface {
	httpHandler(url string, method string) (int, string)
}