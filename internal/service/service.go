package service

import . "digit-caster/internal/model"

type Calculations interface {
	Calculate(calcState CalcState) float32
	addition(firstMember, secondMember float32) float32
	subtraction(firstMember, secondMember float32) float32
	multiplication(firstMember, secondMember float32) float32
	division(firstMember, secondMember float32) float32
	square(firstMember, secondMember float32) float32
}

type Services struct {
	CalculationService Calculations
}

func NewServices() *Services {
	return &Services{
		CalculationService: NewCalculationService(),
	}
}
