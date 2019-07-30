package main

import (
	"math"

	"github.com/unixpickle/num-analysis/linalg"
	"github.com/unixpickle/num-analysis/linalg/leastsquares"
)

func BestFit(basis []PolarFunc, particles []*Particle) []float64 {
	var thetas []float64
	var radii []float64
	for _, p := range particles {
		thetas = append(thetas, math.Atan2(p.Position.Y, p.Position.X))
		radii = append(radii, p.Position.Norm())
	}

	matrix := linalg.NewMatrix(len(particles), len(basis))
	for i, theta := range thetas {
		for j, b := range basis {
			matrix.Set(i, j, b.Eval(theta))
		}
	}

	solver := leastsquares.NewSolver(matrix)
	return solver.Solve(radii)
}

type CombinedFunc struct {
	Basis  []PolarFunc
	Coeffs []float64
}

func (c *CombinedFunc) Eval(theta float64) float64 {
	var res float64
	for i, b := range c.Basis {
		res += c.Coeffs[i] * b.Eval(theta)
	}
	return res
}
