package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	f, err := cubica([]float64{-2, -1, 0, 1, 2}, []float64{3, 0, 2, 4, 4})
	if err != nil {
		panic(err)
	}
	fmt.Println(f(-2))
	fmt.Println(f(-1))
	fmt.Println(f(0))
	fmt.Println(f(1))
	fmt.Println(f(2))
	f2, err := cuadratica([]float64{-2, -1, 0, 1, 2}, []float64{9.7, 5.2, 3.4, 4.5, 8.5})
	if err != nil {
		panic(err)
	}
	fmt.Println(f2(-2))
	fmt.Println(f2(-1))
	fmt.Println(f2(0))
	fmt.Println(f2(1))
	fmt.Println(f2(2))
}

type MatrizExtendida struct {
	ecuaciones [][]float64
	igualdades []float64
}

func (m MatrizExtendida) Mostrar() {
	charr := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i := range len(m.ecuaciones) {
		for j := range len(m.ecuaciones) {
			fmt.Printf("%v%v", m.ecuaciones[i][j], charr[j%26])
			if j < len(m.ecuaciones)-1 {
				fmt.Print(" + ")
			}
		}
		fmt.Printf(" = %v\n", m.igualdades[i])
	}
}

func (m MatrizExtendida) Organizar() {
	for i := range len(m.ecuaciones) {
		if m.ecuaciones[i][i] != 0 {
			continue
		}
		for j := range len(m.ecuaciones) {
			if i == j {
				continue
			}
			if m.ecuaciones[j][i] != 0 {
				temp := m.ecuaciones[i]
				m.ecuaciones[i] = m.ecuaciones[j]
				m.ecuaciones[j] = temp
				temp2 := m.igualdades[i]
				m.igualdades[i] = m.igualdades[j]
				m.igualdades[j] = temp2
				break
			}
		}
	}
}

func (m MatrizExtendida) Montante() []float64 {
	pivote := 1.0
	prev_ecuaciones := make([][]float64, len(m.ecuaciones))
	prev_igualdades := make([]float64, len(m.ecuaciones))
	for i := range len(m.ecuaciones) {
		fmt.Println()
		for j := range len(m.ecuaciones) {
			prev_ecuaciones[j] = make([]float64, len(m.ecuaciones[j]))
			prev_igualdades[j] = m.igualdades[j]
			for k := range len(m.ecuaciones[j]) {
				prev_ecuaciones[j][k] = m.ecuaciones[j][k]
			}
		}
		for row := range len(m.ecuaciones) {
			for col := range len(m.ecuaciones) + 1 {
				if row == col && row == i || row == i {
					continue
				}
				if col == i {
					m.ecuaciones[row][col] = 0
					continue
				}
				if col < len(m.ecuaciones) {
					m.ecuaciones[row][col] = (prev_ecuaciones[i][i]*prev_ecuaciones[row][col] - prev_ecuaciones[row][i]*prev_ecuaciones[i][col]) / pivote
				} else {
					m.igualdades[row] = (prev_ecuaciones[i][i]*prev_igualdades[row] - prev_ecuaciones[row][i]*prev_igualdades[i]) / pivote
				}
			}
		}
		fmt.Printf("iteracion: %v\npivote: %v\n", i+1, pivote)
		pivote = m.ecuaciones[i][i]
		m.Mostrar()
	}
	results := make([]float64, len(m.ecuaciones))
	for i := range m.ecuaciones {
		results[i] = m.igualdades[i] / pivote
	}
	return results
}

func cubica(x []float64, y []float64) (func(float64) float64, error) {
	if len(x) != len(y) {
		return func(x float64) float64 { return 0.0 }, errors.New("lengths don't match")
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
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
	m := new(MatrizExtendida)
	m.ecuaciones = [][]float64{
		{sums[3], sums[2], sums[0], float64(len(x))},
		{sums[4], sums[3], sums[2], sums[0]},
		{sums[5], sums[4], sums[3], sums[2]},
		{sums[6], sums[5], sums[4], sums[3]},
	}
	m.igualdades = []float64{
		sums[1], sums[7], sums[8], sums[9],
	}
	m.Organizar()
	fmt.Println()
	m.Mostrar()
	coeficientes := m.Montante()
	return func(x float64) float64 {
		return coeficientes[0]*math.Pow(x, 3) + coeficientes[1]*math.Pow(x, 2) + coeficientes[2]*x + coeficientes[3]
	}, nil
}

func cuadratica(x []float64, y []float64) (func(float64) float64, error) {
	if len(x) != len(y) {
		return func(x float64) float64 { return 0.0 }, errors.New("lengths don't match")
	}
	sums := [7]float64{}
	fmt.Printf("x\ty\tx2\tx3\tx4\txy\tx2y\n")
	for i := range x {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", x[i], y[i], math.Pow(x[i], 2), math.Pow(x[i], 3), math.Pow(x[i], 4), x[i]*y[i], math.Pow(x[i], 2)*y[i])
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
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
	m := new(MatrizExtendida)
	m.ecuaciones = [][]float64{
		{sums[2], sums[0], float64(len(x))},
		{sums[3], sums[2], sums[0]},
		{sums[4], sums[3], sums[2]},
	}
	m.igualdades = []float64{
		sums[1], sums[5], sums[6],
	}
	m.Organizar()
	fmt.Println()
	m.Mostrar()
	coeficientes := m.Montante()
	return func(x float64) float64 { return coeficientes[0]*math.Pow(x, 2) + coeficientes[1]*x + coeficientes[2] }, nil
}
