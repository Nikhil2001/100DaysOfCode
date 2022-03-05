//print subsequences of an array using Recursion

package main

import "fmt"

var arr1 = []int{}

func subsequence(index int, arr []int) {

	if index == len(arr) {
		fmt.Println(arr1)
		return
	}
	arr1 = append(arr1, arr[index])
	subsequence(index+1, arr)
	arr1 = arr1[:len(arr1)-1]
	subsequence(index+1, arr)

}

func main() {
	arr := []int{3, 1, 2}
	subsequence(0, arr)
}

/*
nikhil@nikhil:~/job/100DaysOfCode/recursion$ go run recrs7.go
[3 1 2]
[3 1]
[3 2]
[3]
[1 2]
[1]
[2]
[]
nikhil@nikhil:~/job/100DaysOfCode/recursion$ */
