package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	x := []float64{1.9, 2.4, 4.8, 5.2}
	y := []float64{2.5, 2.7, 3.7, 5.2}
	f := func(x float64) float64 { return 1 / math.Tan(x) }
	g, err := LinealConFuncion(x, y, f)
	if err != nil {
		panic(err)
	}
	fmt.Println(g(1.9))
	fmt.Println(g(2.4))
	fmt.Println(g(4.8))
	fmt.Println(g(5.2))
	g, err = CuadraticaConFuncion(x, y, f)
	if err != nil {
		panic(err)
	}
	fmt.Println(g(1.9))
	fmt.Println(g(2.4))
	fmt.Println(g(4.8))
	fmt.Println(g(5.2))

	x = []float64{4.8, 5.1, 10.2, 10.9}
	y = []float64{3.7, 5.2, 6.0, 8.3}
	f = func(x float64) float64 { return math.Tan(x) }
	g, err = LinealConFuncion(x, y, f)
	if err != nil {
		panic(err)
	}
	fmt.Println(g(4.8))
	fmt.Println(g(5.1))
	fmt.Println(g(10.2))
	fmt.Println(g(10.9))
	g, err = CuadraticaConFuncion(x, y, f)
	if err != nil {
		panic(err)
	}
	fmt.Println(g(4.8))
	fmt.Println(g(5.1))
	fmt.Println(g(10.2))
	fmt.Println(g(10.9))

	/* f, err := cubica([]float64{-2, -1, 0, 1, 2}, []float64{3, 0, 2, 4, 4})
	if err != nil {
		panic(err)
	}
	fmt.Println(f(-2))
	fmt.Println(f(-1))
	fmt.Println(f(0))
	fmt.Println(f(1))
	fmt.Println(f(2))
	f2, err := cuadratica([]float64{-2, -1, 0, 1, 2}, []float64{12.7, 3.2, -3.3, -6.8, -7.3})
	if err != nil {
		panic(err)
	}
	fmt.Println(f2(-2))
	fmt.Println(f2(-1))
	fmt.Println(f2(0))
	fmt.Println(f2(1))
	fmt.Println(f2(2)) */
}

type MatrizExtendida struct {
	Ecuaciones [][]float64
	Igualdades []float64
}

func (m MatrizExtendida) Mostrar() {
	for i := range len(m.Ecuaciones) {
		for j := range len(m.Ecuaciones) {
			fmt.Printf("%.9f a%v", m.Ecuaciones[i][j], i)
			if j < len(m.Ecuaciones)-1 {
				fmt.Print(" + ")
			}
		}
		fmt.Printf(" = %.9f\n", m.Igualdades[i])
	}
}

func (m MatrizExtendida) Organizar() {
	for i := range len(m.Ecuaciones) {
		if m.Ecuaciones[i][i] != 0 {
			continue
		}
		for j := range len(m.Ecuaciones) {
			if i == j {
				continue
			}
			if m.Ecuaciones[j][i] != 0 {
				temp := m.Ecuaciones[i]
				m.Ecuaciones[i] = m.Ecuaciones[j]
				m.Ecuaciones[j] = temp
				temp2 := m.Igualdades[i]
				m.Igualdades[i] = m.Igualdades[j]
				m.Igualdades[j] = temp2
				break
			}
		}
	}
}

func (m MatrizExtendida) Montante() []float64 {
	pivote := 1.0
	prev_ecuaciones := make([][]float64, len(m.Ecuaciones))
	prev_igualdades := make([]float64, len(m.Ecuaciones))
	for i := range len(m.Ecuaciones) {
		fmt.Println()
		for j := range len(m.Ecuaciones) {
			prev_ecuaciones[j] = make([]float64, len(m.Ecuaciones[j]))
			prev_igualdades[j] = m.Igualdades[j]
			for k := range len(m.Ecuaciones[j]) {
				prev_ecuaciones[j][k] = m.Ecuaciones[j][k]
			}
		}
		for row := range len(m.Ecuaciones) {
			for col := range len(m.Ecuaciones) + 1 {
				if row == col && row == i || row == i {
					continue
				}
				if col == i {
					m.Ecuaciones[row][col] = 0
					continue
				}
				if col < len(m.Ecuaciones) {
					m.Ecuaciones[row][col] = (prev_ecuaciones[i][i]*prev_ecuaciones[row][col] - prev_ecuaciones[row][i]*prev_ecuaciones[i][col]) / pivote
				} else {
					m.Igualdades[row] = (prev_ecuaciones[i][i]*prev_igualdades[row] - prev_ecuaciones[row][i]*prev_igualdades[i]) / pivote
				}
			}
		}
		fmt.Printf("iteracion: %v\npivote: %v\n", i+1, pivote)
		pivote = m.Ecuaciones[i][i]
		m.Mostrar()
	}
	results := make([]float64, len(m.Ecuaciones))
	for i := range m.Ecuaciones {
		results[i] = m.Igualdades[i] / pivote
	}
	return results
}

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
