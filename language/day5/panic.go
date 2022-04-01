package main

import (
	"fmt"
)

func Panic() {
	defer fmt.Println("deffered call -1")
	defer func() {
		fmt.Println("deffered call -2")
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	panic("paniking for example")
	fmt.Println(" this will be nerver called")
}

func main() {

	fmt.Println("starting to panic")
	Panic()
	fmt.Println("Program regains control after panic recovery")

}
