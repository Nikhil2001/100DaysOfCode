package main

import (
	"fmt"
	"reflect"
)

func main() {

	x := 5
	fmt.Println(reflect.TypeOf(x))

	y := float64(x)
	fmt.Println(reflect.TypeOf(y))

	var i interface{} = 5
	var j = i.(int)
	fmt.Println(j) 
	
//	i  = true
//	var k = i.(int) //results in -> panic: interface conversion: interface {} is bool, not int
//	fmt.Println(k)

}
