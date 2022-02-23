package main

import "fmt"

func add(a,b int) int {
	return a+b 
}

var sub =  func(a,b int) int {
	return a-b 
}

func main(){
	var Op = add
	fmt.Println(Op(1,2))
	Op = sub
	fmt.Println(Op(1,2))

	Op = func(a,b int) int {
		return a*b
	}
	fmt.Println(Op(3,4))
}