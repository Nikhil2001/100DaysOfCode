package main

import (
	"fmt"
	"math"
)

func sum(a, b int) int {
	return a + b
}

func addAndSub(a, b int) (add, sub int) {
	add = a + b
	sub = a - b
	return
}

func addSubMul(a, b int) (int, int, int) {
	return a + b, a - b, a * b
}

func PrintPi() {
	fmt.Printf("%T %v\n", math.Pi, math.Pi)
}

func fib(n int) {
	a, b, c := 0, 1, 0
	fmt.Print("fib:[ ", a, b, " ")
	for i := 2; i < n; i++ {
		c = a + b
		fmt.Print(c, " ")
		a, b = b, c

	}
	fmt.Println("]")
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(addAndSub(1, 2))
	fmt.Println(addSubMul(2, 4))
	PrintPi()
	fib(10)
}
