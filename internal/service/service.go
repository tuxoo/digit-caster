package service

import "context"

type Calculations interface {
	Summation(ctx context.Context)
	Subtraction(ctx context.Context)
	Multiple(ctx context.Context)
	Division(ctx context.Context)
	Square(ctx context.Context)
}

type Services struct {
	CalculationService Calculations
}

func NewServices() *Services {
	return &Services{
		CalculationService: NewCalculationService(),
	}
}
