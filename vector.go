package main

import (
	"math"
	"math/rand"
)

type Vector2 struct {
	X float64
	Y float64
}

func RandomCircleVector2() Vector2 {
	radius := math.Sqrt(rand.Float64())
	theta := rand.Float64() * math.Pi * 2
	return Vector2{X: math.Cos(theta) * radius, Y: math.Sin(theta) * radius}
}

func (v Vector2) Scale(s float64) Vector2 {
	return Vector2{X: v.X * s, Y: v.Y * s}
}

func (v Vector2) Add(v1 Vector2) Vector2 {
	return Vector2{X: v.X + v1.X, Y: v.Y + v1.Y}
}

func (v Vector2) Sub(v1 Vector2) Vector2 {
	return v.Add(v1.Scale(-1))
}

func (v Vector2) Norm() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v Vector2) OrthogonalVector() Vector2 {
	return Vector2{X: -v.Y, Y: v.X}
}

func (v Vector2) Dot(v1 Vector2) float64 {
	return v.X*v1.X + v.Y*v1.Y
}
