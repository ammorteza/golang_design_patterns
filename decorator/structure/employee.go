package structure

import "math"

type Employee struct {
	Name 		string
	BaseIncome 	uint
}

func NewEmployee(name string, baseIncome uint) GovEmployee {
	return &Employee{
		Name: name,
		BaseIncome: baseIncome,
	}
}

func (this *Employee)Income() uint{
	return this.BaseIncome + uint(math.Round(float64(this.BaseIncome) * 0.25))
}