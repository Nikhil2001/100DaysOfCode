package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)

	arr[2] = 100
	fmt.Println(arr, len(arr))

	arr1 := [...]string{"nikhil", "nikhitha"}
	fmt.Println(arr1, len(arr1))

}
