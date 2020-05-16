package main

type IShape interface {
	SetName(string)
	Name() string
}