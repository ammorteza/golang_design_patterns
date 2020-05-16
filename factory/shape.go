package main

type Shape struct {
	name 		string
}

func (s *Shape)Name() string{
	return s.name
}

func (s *Shape)SetName(name string){
	s.name = name
}