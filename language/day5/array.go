package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)

	arr[2] = 100
	fmt.Println(arr, len(arr))

	arr1 := [...]string{5: "nikhil", 1: "nikhitha"}
	fmt.Println(arr1, len(arr1))

	for k, v := range arr1 {
		fmt.Println(k, v)
	}

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

}
