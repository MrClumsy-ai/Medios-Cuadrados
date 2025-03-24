package main

import (
	"errors"
	"fmt"
	"math"
)

// g(x) = a + bx + ax^2 + a3f(x)
func CuadraticaConFuncion(x []float64, y []float64, f func(float64) float64) (func(float64) float64, error) {
	fmt.Println("##########################################################################################################################")
	if len(x) != len(y) {
		return nil, errors.New("lengths dont match")
	}
	sums := [12]float64{}
	fmt.Printf("x\t\ty\t\tx2\t\tx3\t\tx4\t\tf(x)\t\txf(x)\t\tx2f(x)\t\tf(x)2\t\txy\t\tx2y\t\tyf(x)\n")
	for i := range x {
		fmt.Printf("%.9f\t%.9f\t%.9f\t%.9f\t%.9f\t%.9f\t%.9f\t%.9f\t%.9f\t%.9f\t%.9f\t%.9f\n",
			x[i], y[i], math.Pow(x[i], 2), math.Pow(x[i], 3), math.Pow(x[i], 4), f(x[i]), x[i]*f(x[i]), math.Pow(x[i], 2)*f(x[i]), math.Pow(f(x[i]), 2), x[i]*y[i], math.Pow(x[i], 2)*y[i], y[i]*f(x[i]),
		)
		sums[0] += x[i]
		sums[1] += y[i]
		sums[2] += math.Pow(x[i], 2)
		sums[3] += math.Pow(x[i], 3)
		sums[4] += math.Pow(x[i], 4)
		sums[5] += f(x[i])
		sums[6] += x[i] * f(x[i])
		sums[7] += math.Pow(x[i], 2) * f(x[i])
		sums[8] += math.Pow(f(x[i]), 2)
		sums[9] += x[i] * y[i]
		sums[10] += math.Pow(x[i], 2) * y[i]
		sums[11] += y[i] * f(x[i])
	}
	fmt.Println("sumas:")
	for _, v := range sums {
		fmt.Printf("%.9f\t", v)
	}
	fmt.Println()
	m := new(MatrizExtendida)
	m.Ecuaciones = [][]float64{
		{float64(len(x)), sums[0], sums[2], sums[5]},
		{sums[0], sums[2], sums[3], sums[6]},
		{sums[2], sums[3], sums[4], sums[7]},
		{sums[5], sums[6], sums[7], sums[8]},
	}
	m.Igualdades = []float64{
		sums[1], sums[9], sums[10], sums[11],
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
	return func(x float64) float64 {
		return results[0] + results[1]*x + results[2]*math.Pow(x, 2) + results[3]*f(x)
	}, nil
}
