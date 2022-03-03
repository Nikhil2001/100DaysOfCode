//sum of 1 to N numbers using Recursion
package main

import "fmt"

func sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum(n-1)
}

func add(n, sum int) {
	if n == 0 {
		fmt.Println(sum)
		return
	}
	add(n-1, sum+n)

}

func main() {
	n := 5
	fmt.Println(sum(n))
	var sum int
	add(n, sum)
}
