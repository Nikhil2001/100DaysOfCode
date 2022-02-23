package main

import "fmt"


var(
	mul = func(a,b int) int {
		return a*b
	}

	square = func(a int) int {
		return mul(a,a)
	}
)


func main(){
	fmt.Println(mul(3,4))
	fmt.Println(square(5))
}