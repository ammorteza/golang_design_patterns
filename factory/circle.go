package main

type Circle struct {
	Shape
}

func NewCircle(name string) IShape{
	circle := Circle{}
	circle.SetName(name)
	return &circle
}