package main

import "fmt"

func Same[T comparable](a, b T) bool {
	return a == b
}

func Print[T any](a T) T {
	return a
}

func main() {
	fmt.Println(Same(1, 1))
	fmt.Println(Same(1.4, 1.0))
	fmt.Println(Same(1.4+3i, 1.0+3i))
	fmt.Println(Same("nikhil", "nikhil"))
	fmt.Println(Same(true, false))
	fmt.Println(Print(true))
	fmt.Println(Print(1))
}

