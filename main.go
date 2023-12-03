package main

import (
	"fmt"
	"math"
)

func computeRungeKuta4(f func(float64, float64) float64, initial, a, n int, h float64) [][]float64 {
	xn := float64(a)
	yn := float64(initial)
	resultTable := [][]float64{{xn, yn}}
	currentN := 1
	fmt.Println(xn, yn)

	for currentN < n {
		tempX := xn
		tempY := yn
		k1 := f(tempX, tempY) * h
		tempX = xn + h/2
		tempY = yn + k1/2
		k2 := f(tempX, tempY) * h
		tempX = xn + h/2
		tempY = yn + k2/2
		k3 := f(tempX, tempY) * h
		tempX = xn + h
		tempY = yn + k3
		k4 := f(tempX, tempY) * h
		yn = yn + (1.0/6)*(k1+2*k2+2*k3+k4)
		xn = xn + h
		form := fmt.Sprintf("%.3f, %.3f, %.3f, %.3f, %.3f, %.3f", k1, k2, k3, k4, xn, yn)
		fmt.Println(form)
		resultTable = append(resultTable, []float64{xn, yn})
		currentN++
	}

	return resultTable
}

func extrapolateAdams4(y0, yk, yk1, yk2, yk3 float64, h float64) float64 {
	return y0 + (h/24)*(55*yk-59*yk1+37*yk2-9*yk3)
}

func adams(valueTable [][]float64, h float64, f func(float64, float64) float64, n int) [][]float64 {
	resultTable := [][]float64{}
	for _, value := range valueTable {
		resultTable = append(resultTable, []float64{value[0], value[1], f(value[0], value[1])})
	}

	i := 0
	for i < n {
		xVal := resultTable[len(resultTable)-1][0]
		xVal += h
		yVal := extrapolateAdams4(resultTable[len(resultTable)-1][1], resultTable[len(resultTable)-1][2],
			resultTable[len(resultTable)-2][2], resultTable[len(resultTable)-3][2], resultTable[len(resultTable)-4][2], h)
		value := []float64{xVal, yVal, f(xVal, yVal)}
		resultTable = append(resultTable, value)
		i++
	}

	return resultTable
}

func computeRungeKuta2(f func(float64, float64) float64, initial, u, a, n int, h float64) [][]float64 {
	un := float64(u)
	xn := float64(a)
	yn := float64(initial)
	resultTable := [][]float64{{xn, yn}}
	currentN := 1
	fmt.Println(xn, yn)

	for currentN < n {
		tempX := xn
		tempY := yn
		tempU := un
		k1y := tempU * h
		k1 := f(tempX, tempY) * h
		tempX = xn + 3.0/4*h
		tempY = yn + +3.0/4*h
		k2 := f(tempX, tempY) * h
		un = un + 1.0/3*k1 + 2.0/3*k2
		yn = yn + k1y
		xn = xn + h
		form := fmt.Sprintf("%.3f, %.3f, %.3f, %.3f", k1, k2, xn, yn)
		fmt.Println(form)
		resultTable = append(resultTable, []float64{xn, yn})
		currentN++
	}

	return resultTable
}

func main() {

	f := func(x, y float64) float64 {
		return -math.Pow(y-x, 0.5) / 2.0 * x
	}
	res := computeRungeKuta2(f, 2, 1, 1, 11, 0.1)

	for _, value := range res {
		fmt.Printf("%.3f, %.3f\n", value[0], value[1])
	}
}
