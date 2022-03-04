//Print Series
// if n = 5, k =2 Print [5,3,1,-1,1,3,5] using recursion
package main

import "fmt"

func Print(output *[]int, n, k int) {
	*output = append(*output, n)
	if n < 0 {
		return
	}
	Print(output, n-k, k)
	*output = append(*output, n)
}

func main() {
	var output = []int{}

	n, k := 5, 2
	Print(&output, n, k)
	fmt.Println(output)

}
