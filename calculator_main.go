package main

import (
	"fmt"
	"math"
)

type inpout struct {
	number1 float64
	number2 float64
}

func (inp *inpout) sum() float64 {
	var sum float64
	sum = inp.number1 + inp.number2
	return sum
}
func (inp *inpout) Minus() float64 {
	var minus float64
	minus = inp.number1 - inp.number2
	return minus
}
func (inp *inpout) Multi() float64 {
	var multi float64
	multi = inp.number1 * inp.number2
	return multi
}
func (inp *inpout) Divided() float64 {
	var divid float64
	divid = inp.number1 / inp.number2
	return divid
}
func (inp *inpout) Power() float64 {
	var powr float64
	powr = math.Pow(inp.number1, inp.number2)
	return powr
}
func (inp *inpout) square() float64 {
	var sqr float64
	sqr = math.Sqrt(inp.number1)
	return sqr
}

func main() {
	var number1 float64
	var number2 float64
	var operator string

	fmt.Printf("First Number:")
	fmt.Scanln(&number1)
	fmt.Printf("operator:")
	fmt.Scanln(&operator)
	fmt.Printf("Second Number:")
	fmt.Scanln(&number2)

	inpouting := inpout{
		number1: number1,
		number2: number2,
	}

	switch operator {
	case "+":
		fmt.Println(inpouting.sum())
	case "-":
		fmt.Println(inpouting.Minus())
	case "*":
		fmt.Println(inpouting.Multi())
	case "/":
		fmt.Println(inpouting.Divided())
	case "^":
		fmt.Println(inpouting.Power())
	case "√":
		fmt.Println(inpouting.square())
	case "C", "c":
		break
	default:
		fmt.Println("it can't support")
	}
}
