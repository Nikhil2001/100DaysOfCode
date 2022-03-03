//reverse an array using Recursion
//check string is palindrome using Recursion
package main

import "fmt"

//var arr = [...]int{1, 2, 3, 4, 5}
var arr = [...]int{1, 2, 3, 4, 5, 6}

func ReverseArray(i, j int) {
	if i >= len(arr)/2 {
		return
	}
	c := arr[i]
	arr[i] = arr[j]
	arr[j] = c
	ReverseArray(i+1, j-1)

}

var str = "MADAM"
var strRunes = []rune(str)

func Palindrome(i int) bool {
	if i >= len(strRunes)/2 {
		return true
	}

	if strRunes[i] != strRunes[len(strRunes)-i-1] {
		return false
	}
	return Palindrome(i + 1)

}

func main() {
	fmt.Println(arr)
	ReverseArray(0, len(arr)-1)
	fmt.Println(arr)

	fmt.Println(strRunes)

	fmt.Println(Palindrome(0))

}
