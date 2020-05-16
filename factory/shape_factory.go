package main

import "errors"

func GetShape(_type, name string) (IShape, error){
	switch _type {
	case "circle":
		return NewCircle(name), nil
	case "square":
		return NewSquare(name), nil
	default:
		return nil, errors.New("current factory cannot produces " + name)
	}
}