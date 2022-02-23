package main

import "fmt"

func avg(nums ...float64) float64 {
	n := len(nums)
	sum := 0.0
	for _, v := range nums {
		sum = sum + v
	}
	return sum / float64(n)
}

func main() {
	fmt.Println(avg(1.0, 2.0))
	fmt.Println(avg(1.0, 2.0, 3.0))
	nums := []float64{1.0, 2.0, 3.6}
	fmt.Printf("%.1f\n",avg(nums...))
}
