package main

import "fmt"

var count int

func name() {
	fmt.Println("Nikhil")
	count++

	if count == 5 {
		return
	}
	name()
}

func PrintNum() {
	count++
	fmt.Println(count)
	if count == 5 {
		return
	}
	PrintNum()
}

func NumbersBack() {
	count++
	if count == 5 {
		fmt.Println(count)
		return
	}
	NumbersBack()
	count--
	fmt.Println(count)
}

var num = 1
func Numbers() {
	count++
	if count == 5 {
		fmt.Println(num)
		return
	}
	Numbers()
	num++
	fmt.Println(num)
}


func main() {
	name()
	
	count = 0
	PrintNum()

	count = 0
	NumbersBack()

	count = 0
	Numbers()
}
