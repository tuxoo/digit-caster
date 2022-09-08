package model

type CalcState struct {
	PreviousNum float64 `json:"previousNum"`
	CurrentNum  float64 `json:"currentNum"`
	Operation   string  `json:"operation"`
}
