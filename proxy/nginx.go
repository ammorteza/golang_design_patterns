package main

import "net/http"

type Nginx struct {
	router 					*Router
	maxAllowedRequest 		int
	requestCountPerUrl		map[string]int
}

func NewNginx() Server{
	return &Nginx{
		router: NewRouter(),
		maxAllowedRequest: 2,
		requestCountPerUrl: make(map[string]int),
	}
}

func (n *Nginx)httpHandler(url string, method string) (int, string){
	if n.checkRequestHistory(url, method) {
		n.requestCountPerUrl[url + method] = n.requestCountPerUrl[url + method] + 1
		return n.router.httpHandler(url, method)
	}else{
		return http.StatusNotFound, "not found"
	}
}

func (n *Nginx)checkRequestHistory(url string, method string) bool{
	count, ok := n.requestCountPerUrl[url + method]
	if !ok || count < n.maxAllowedRequest{
		return true
	}
	return false
}