package main

import "fmt"

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
