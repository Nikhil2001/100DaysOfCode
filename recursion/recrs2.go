package main

import "fmt"

var count int = 0

func f() {
	fmt.Println(1)
    count++
	if count == 10{
		return // base condition
	}
	f()
}


func main(){
	f()
}