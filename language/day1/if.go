package main

import "fmt"

func main() {
	c := true

	if c {
		fmt.Println("c is true")
	}

	num := 31
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	a := 0
	if a > 0 {
		fmt.Println(a, "is positive")
	} else if a < 0 {
		fmt.Println(a, "is negative")
	} else {
		fmt.Println("a is", 0)
	}

	if dirx := "North"; dirx == "East" {
		fmt.Println("direction is East")
	} else if dirx == "west" {
		fmt.Println("direction is West")
	} else if dirx == "South" {
		fmt.Println("direction is South")
	} else {
		fmt.Println("direction is North")
	}

}

/* OUTPUT
nikhil@nikhil:~/job/100DaysOfCode/language/day1$ go run if.go
c is true
31 is odd
a is 0
direction is North
*/
