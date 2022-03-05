//Program to print digits using recursion
package main

import "fmt"

func saydigits(n int) {

	if n == 0 {
		return
	}

	saydigits(n / 10)
	fmt.Print(n%10, " ")
}

func main() {
	saydigits(43921)
	fmt.Println()
}

/*
nikhil@nikhil:~/job/100DaysOfCode$ go run recursion/recrs12.go
4 3 9 2 1
*/
