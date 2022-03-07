package main

import "fmt"

func BinarySearch(arr []int, num int) (index int, found bool) {

	start := 0
	end := len(arr) - 1

	for {
		if start == end {
			if arr[start] == num {
				return start, true
			}
			return -1, false
		}
		mid := (start + end) / 2

		switch {
		case arr[mid] == num:
			return mid, true
		case arr[mid] < num:
			start = mid + 1
		case arr[mid] > num:
			end = mid - 1
		}

	}

}

func main() {
	arr := []int{1, 2, 3, 5, 6, 7}
	fmt.Println(BinarySearch(arr, 99))
}
