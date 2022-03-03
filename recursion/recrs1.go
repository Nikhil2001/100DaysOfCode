package main

import "fmt"

func f() {
	fmt.Println(1)
	//recursive calling
	f()
	//running will prints 1 untill  stack overflows
}


func main(){
	f()
}