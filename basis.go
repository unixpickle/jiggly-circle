package main

import "math"

type PolarFunc interface {
	Eval(theta float64) float64
}

func StandardBasis() []PolarFunc {
	funcs := make([]PolarFunc, 1, 15)
	funcs[0] = UnitFunc{}
	for i := 1; i <= 7; i++ {
		funcs = append(
			funcs,
			&TrigPolarFunc{Cos: false, Frequency: i},
			&TrigPolarFunc{Cos: true, Frequency: i},
		)
	}
	return funcs
}

type TrigPolarFunc struct {
	Cos       bool
	Frequency int
}

func (t *TrigPolarFunc) Eval(theta float64) float64 {
	if t.Cos {
		return math.Cos(theta * float64(t.Frequency))
	} else {
		return math.Sin(theta * float64(t.Frequency))
	}
}

type UnitFunc struct{}

func (u UnitFunc) Eval(theta float64) float64 {
	return 1
}
