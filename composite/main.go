package main

import "fmt"

type Employee interface {
	Print()
}

type Developer struct {
	Name 	string
}

type Manager struct {
	Employees 	[]Employee
	Name 		string
}

type Office struct {
	Employees 	[]Employee
	Name 		string
}

func NewDeveloper(name string) *Developer{
	return &Developer{
		Name: name,
	}
}

func NewManager(name string, emps ...Employee) *Manager{
	return &Manager{
		Name: name,
		Employees: emps,
	}
}

func NewOffice(name string, emps ...Employee) *Office{
	return &Office{
		Name: name,
		Employees: emps,
	}
}

func (this *Developer)Print(){
	fmt.Println("/////////////// " , this.Name , " ///////////////")
	fmt.Println(this.Name)
	fmt.Println("/////////////// end of " , this.Name , " ///////////////")
}

func (this *Manager)Print(){
	fmt.Println("/////////////// " , this.Name , " ///////////////")
	for _, emp := range this.Employees{
		emp.Print()
	}
	fmt.Println("/////////////// end of " , this.Name , " ///////////////")
}

func (this *Office)Print(){
	fmt.Println("/////////////// " , this.Name , " ///////////////")
	for _, emp := range this.Employees{
		emp.Print()
	}
	fmt.Println("/////////////// end of " , this.Name , " ///////////////")
}

func main(){
	dev1 := NewDeveloper("Developer 1")
	dev2 := NewDeveloper("Developer 2")
	dev3 := NewDeveloper("Developer 3")
	dev4 := NewDeveloper("Developer 4")

	mng1 := NewManager("Manager 1", dev1, dev2)
	mng2 := NewManager("Manager 2", mng1, dev3)

	office := NewOffice("Snap!", mng2, dev4)

	office.Print()
}