package main

import (
	"errors"
	"fmt"
	"math"
)

func Cubica(x []float64, y []float64) (func(float64) float64, error) {
	if len(x) != len(y) {
		return nil, errors.New("lengths don't match")
	}
	sums := [10]float64{}
	fmt.Printf("x\ty\tx2\tx3\tx4\tx5\tx6\txy\tx2y\tx3y\n")
	for i := range x {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", x[i], y[i], math.Pow(x[i], 2), math.Pow(x[i], 3), math.Pow(x[i], 4), math.Pow(x[i], 5), math.Pow(x[i], 6), x[i]*y[i], math.Pow(x[i], 2)*y[i], math.Pow(x[i], 3)*y[i])
		sums[0] += x[i]
		sums[1] += y[i]
		sums[2] += math.Pow(x[i], 2)
		sums[3] += math.Pow(x[i], 3)
		sums[4] += math.Pow(x[i], 4)
		sums[5] += math.Pow(x[i], 5)
		sums[6] += math.Pow(x[i], 6)
		sums[7] += x[i] * y[i]
		sums[8] += math.Pow(x[i], 2) * y[i]
		sums[9] += math.Pow(x[i], 3) * y[i]
	}
	fmt.Println("sumas:")
	for _, v := range sums {
		fmt.Printf("%.9f\t", v)
	}
	fmt.Println()
	m := new(MatrizExtendida)
	m.Ecuaciones = [][]float64{
		{sums[3], sums[2], sums[0], float64(len(x))},
		{sums[4], sums[3], sums[2], sums[0]},
		{sums[5], sums[4], sums[3], sums[2]},
		{sums[6], sums[5], sums[4], sums[3]},
	}
	m.Igualdades = []float64{
		sums[1], sums[7], sums[8], sums[9],
	}
	m.Organizar()
	fmt.Println()
	m.Mostrar()
	coeficientes := m.Montante()
	fmt.Println(coeficientes)
	return func(x float64) float64 {
		return coeficientes[0]*math.Pow(x, 3) + coeficientes[1]*math.Pow(x, 2) + coeficientes[2]*x + coeficientes[3]
	}, nil
}
