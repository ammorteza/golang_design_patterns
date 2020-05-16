package main

type Square struct {
	Shape
}

func NewSquare(name string) IShape{
	square := Square{}
	square.SetName(name)
	return &square
}