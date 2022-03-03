//Print Name n times using Recursion
//Print 1 to n  using Recursion
//Print n to 1  using Recursion

package main

import "fmt"

func PrintName(n int) {
	if n == 0 {
		return
	}
	fmt.Println("Nikhil")
	PrintName(n-1)
}


func PrintNumbers(i,n int){
	if i > n {
		return
	} 
	fmt.Println(i)
	PrintNumbers(i+1,n)
}


func PrintNumbersReverse(i,n int){
	if i > n {
		return
	} 	
	PrintNumbersReverse(i+1,n)
	fmt.Println(i)
}

func main() {
	n := 5
	PrintName(5)
	PrintNumbers(1,n)
	fmt.Println("Numbers Reverse")
	PrintNumbersReverse(1,n)
}
