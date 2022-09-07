package model

type CalcState struct {
	PreviousNum float32 `json:"previousNum"`
	CurrentNum  float32 `json:"currentNum"`
	Operation   string  `json:"operation"`
}
