package main

import (
	"fmt"
	"github.com/ammorteza/design_patterns/function"
	"github.com/ammorteza/design_patterns/structure"
)

func main(){
	fmt.Println("///////////////function//////////////////")
	income := function.Income
	withInsurance := function.WithInsurance(income, 0.07)
	withTax := function.WithTax(withInsurance, 0.15)
	fmt.Println(income(2500000))
	fmt.Println(withInsurance(2500000))
	fmt.Println(withTax(2500000))

	fmt.Println("///////////////structure//////////////////")
	emp := structure.NewEmployee("morteza amzajerdi", 2500000)
	remp := structure.NewREmployee(emp.(*structure.Employee), "23659254")
	gemp := structure.NewGEmployee(emp.(*structure.Employee))
	fmt.Println(emp.Income())
	fmt.Println(remp.Income())
	fmt.Println(gemp.Income())
}