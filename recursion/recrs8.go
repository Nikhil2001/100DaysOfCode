package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)

}

func fibonocci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return factorial(n-1) + factorial(n-2)

}

func main() {
	var n int
	fmt.Println("factorial of ?")
	fmt.Scan(&n)
	fmt.Println(factorial(n))

}
