package service

import (
	. "digit-caster/internal/model"
	"math"
	"strconv"
)

type CalculationService struct {
}

func NewCalculationService() *CalculationService {
	return &CalculationService{}
}

func (s *CalculationService) Calculate(calcState CalcState) string {
	var result float64
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
		result = s.square(calcState.PreviousNum, calcState.CurrentNum)
	}
	return strconv.FormatFloat(result, 'f', -1, 32)
}

func (s *CalculationService) addition(firstMember, secondMember float64) float64 {
	return firstMember + secondMember
}

func (s *CalculationService) subtraction(firstMember, secondMember float64) float64 {
	return firstMember - secondMember
}

func (s *CalculationService) multiplication(firstMember, secondMember float64) float64 {
	return firstMember * secondMember
}

func (s *CalculationService) division(firstMember, secondMember float64) float64 {
	return firstMember / secondMember
}

func (s *CalculationService) square(firstMember, secondMember float64) float64 {
	return math.Pow(firstMember, secondMember)
}
