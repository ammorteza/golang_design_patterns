package function

import "math"

type IncomeType func(uint) uint

func Income(baseAmount uint) uint{
	return baseAmount + uint(math.Round(float64(baseAmount) * float64(0.25)))
}

func WithInsurance(income IncomeType, percent float64) IncomeType{
	return func(u uint) uint {
		baseIncome := income(u)
		return baseIncome - uint(math.Round(float64(baseIncome) * percent))
	}
}

func WithTax(income IncomeType, percent float64) IncomeType{
	return func(u uint) uint {
		baseIncome := income(u)
		return baseIncome - uint(math.Round(float64(baseIncome) * percent))
	}
}