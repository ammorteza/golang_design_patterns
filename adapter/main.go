package main

func main(){
	me := NewPerson("morteza")

	me.TurnOnDevice(&TV{})
	me.TurnOnDevice(&RadioAdapter{Radio{}})
}