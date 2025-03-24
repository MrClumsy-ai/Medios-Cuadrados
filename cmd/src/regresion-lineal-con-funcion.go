package main

import (
	"errors"
	"fmt"
	"math"
)

// g(x) = a + bx + cf(x)
func LinealConFuncion(x []float64, y []float64, f func(float64) float64) (func(float64) float64, error) {
	fmt.Println("##########################################################################################################################")
	if len(x) != len(y) {
		return nil, errors.New("lengths dont match")
	}
	sums := [8]float64{}
	fmt.Printf("x\t\ty\t\tx2\t\tf(x)\t\txf(x)\t\tf(x)2\t\txy\t\tyf(x)\n")
	for i := range x {
		fmt.Printf("%.9f\t%.9f\t%.9f\t%.9f\t%.9f\t%.9f\t%.9f\t%.9f\n", x[i], y[i], math.Pow(x[i], 2), f(x[i]), x[i]*f(x[i]), math.Pow(f(x[i]), 2), x[i]*y[i], y[i]*f(x[i]))
		sums[0] += x[i]
		sums[1] += y[i]
		sums[2] += math.Pow(x[i], 2)
		sums[3] += f(x[i])
		sums[4] += x[i] * f(x[i])
		sums[5] += math.Pow(f(x[i]), 2)
		sums[6] += x[i] * y[i]
		sums[7] += y[i] * f(x[i])
	}
	fmt.Println("sumas:")
	for _, v := range sums {
		fmt.Printf("%.9f\t", v)
	}
	fmt.Println()
	m := new(MatrizExtendida)
	m.Ecuaciones = [][]float64{
		{float64(len(x)), sums[0], sums[3]},
		{sums[0], sums[2], sums[4]},
		{sums[3], sums[4], sums[5]},
	}
	m.Igualdades = []float64{
		sums[1], sums[6], sums[7],
	}
	m.Organizar()
	fmt.Println()
	m.Mostrar()
	results := m.Montante()
	fmt.Println()
	for i := range results {
		fmt.Printf("a%v = %.9f\n", i, results[i])
	}
	fmt.Println()
	return func(x float64) float64 { return results[0] + results[1]*x + results[2]*f(x) }, nil
}
