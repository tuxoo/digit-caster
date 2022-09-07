package service

import . "digit-caster/internal/model"

type CalculationService struct {
}

func NewCalculationService() *CalculationService {
	return &CalculationService{}
}

func (s *CalculationService) Calculate(calcState CalcState) float32 {
	var result float32
	switch calcState.Operation {
	case "+":
		result = s.addition(calcState.PreviousNum, calcState.CurrentNum)
	case "-":
		result = s.subtraction(calcState.PreviousNum, calcState.CurrentNum)
	case "*":
		result = s.multiplication(calcState.PreviousNum, calcState.CurrentNum)
	case "/":
		result = s.division(calcState.PreviousNum, calcState.CurrentNum)
	case "^":
		result = s.square(calcState.CurrentNum)
	}

	return result
}

func (s *CalculationService) addition(firstMember, secondMember float32) float32 {
	return firstMember + secondMember
}

func (s *CalculationService) subtraction(firstMember, secondMember float32) float32 {
	return firstMember - secondMember
}

func (s *CalculationService) multiplication(firstMember, secondMember float32) float32 {
	return firstMember * secondMember
}

func (s *CalculationService) division(firstMember, secondMember float32) float32 {
	return firstMember / secondMember
}

func (s *CalculationService) square(member float32) float32 {
	return member * member
}
