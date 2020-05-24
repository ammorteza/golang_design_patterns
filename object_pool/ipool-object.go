package main

type iPoolObject interface {
	GetId() int
	GetConnection() string
}