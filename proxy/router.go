package main

import "net/http"

type Router struct {}

func NewRouter() *Router{
	return &Router{}
}

func (r *Router)httpHandler(url string, method string) (int, string){
	return makeHandler(url, method)()
}

func makeHandler(url string, method string) func() (int, string){
	switch {
	case url == "/users" &&  method == "GET":
		return func() (int, string) {
			return http.StatusOK, `[{"name" : "morteza", "family" : "amzajerdi"}, {"name" : "ali", "family": "jahanpak"}]`
		}
	case url == "/user/register" && method == "POST":
		return func() (int, string) {
			return http.StatusOK, "new user was registered"
		}
	default:
		return func() (int, string) {
			return http.StatusNotFound, "not found!"
		}
	}
}