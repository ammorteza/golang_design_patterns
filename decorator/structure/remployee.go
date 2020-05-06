package structure

import "math"

type REmployee struct {
	Employee
	Code 		string
}

func NewREmployee(employee *Employee, code string) GovEmployee{
	return &REmployee{
		Employee: *employee,
		Code: code,
	}
}

func (this *REmployee)Income() uint{
	emIncome := this.Employee.Income()
	temp := emIncome
	emIncome -= uint(math.Round(float64(temp) * 0.07)) //insurance percent
	emIncome -= uint(math.Round(float64(temp) * 0.09)) //tax percent
	return emIncome
}