package structure

import "math"

type GEmployee struct {
	Employee
}

func NewGEmployee(employee *Employee) GovEmployee{
	return &GEmployee{
		Employee: *employee,
	}
}

func (this *GEmployee)Income() uint{
	baseIncome := this.Employee.Income()
	temp := baseIncome
	baseIncome -= uint(math.Round(float64(temp) * 0.06)) // insurance percent
	baseIncome -= uint(math.Round(float64(temp) * 0.07)) // tax percent
	return baseIncome
}