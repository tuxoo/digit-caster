package service

import . "digit-caster/internal/model"

type Calculations interface {
	Calculate(calcState CalcState) string
	addition(firstMember, secondMember float64) float64
	subtraction(firstMember, secondMember float64) float64
	multiplication(firstMember, secondMember float64) float64
	division(firstMember, secondMember float64) float64
	square(firstMember, secondMember float64) float64
}

type Services struct {
	CalculationService Calculations
}

func NewServices() *Services {
	return &Services{
		CalculationService: NewCalculationService(),
	}
}
