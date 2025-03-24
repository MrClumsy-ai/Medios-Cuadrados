package main

import (
	"errors"
	"fmt"
	"math"
)

func Cuadratica(x []float64, y []float64) (func(float64) float64, error) {
	if len(x) != len(y) {
		return func(x float64) float64 { return 0.0 }, errors.New("lengths don't match")
	}
	sums := [7]float64{}
	fmt.Printf("x\ty\tx2\tx3\tx4\txy\tx2y\n")
	for i := range x {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\n", x[i], y[i], math.Pow(x[i], 2), math.Pow(x[i], 3), math.Pow(x[i], 4), x[i]*y[i], math.Pow(x[i], 2)*y[i])
		sums[0] += x[i]
		sums[1] += y[i]
		sums[2] += math.Pow(x[i], 2)
		sums[3] += math.Pow(x[i], 3)
		sums[4] += math.Pow(x[i], 4)
		sums[5] += x[i] * y[i]
		sums[6] += math.Pow(x[i], 2) * y[i]
	}
	fmt.Println("sumas:")
	for _, v := range sums {
		fmt.Printf("%.9f\t", v)
	}
	fmt.Println()
	m := new(MatrizExtendida)
	m.Ecuaciones = [][]float64{
		{sums[2], sums[0], float64(len(x))},
		{sums[3], sums[2], sums[0]},
		{sums[4], sums[3], sums[2]},
	}
	m.Igualdades = []float64{
		sums[1], sums[5], sums[6],
	}
	m.Mostrar()
	m.Organizar()
	fmt.Println()
	m.Mostrar()
	coeficientes := m.Montante()
	fmt.Println(coeficientes)
	return func(x float64) float64 { return coeficientes[0]*math.Pow(x, 2) + coeficientes[1]*x + coeficientes[2] }, nil
}
