package main

import "fmt"

const (
	APPLE = iota
	BANANA
	MANGO
	KIWI
)


func main(){
	fmt.Println(KIWI)
	fmt.Printf("%T\n",KIWI)
}