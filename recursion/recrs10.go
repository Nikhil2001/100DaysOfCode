//Print Nth term in GP Series a, ar, ar.pow2 .....ar
package main

import "fmt"

func nthGpTerm(a, r, n int) int {

	if n == 1 {
		return a
	}

	return r * nthGpTerm(a, r, n-1)

}

func Series(a, r, i, n int) {
	if i > n {
		fmt.Printf("\n")
		return
	}
	fmt.Print(nthGpTerm(a, r, i), " ")
	Series(a, r, i+1, n)
}

func main() {
	Series(3, 4, 1, 5)
}
