package main

import (
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
