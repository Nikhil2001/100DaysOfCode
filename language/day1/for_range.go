package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5}

	i := 0
	for range a {
		i++
	}

	fmt.Println(i)

	for i := range a {
		fmt.Println("index",i)
	}

	for i,v := range a {
		fmt.Println("index",i,"value",v)
	}

	for _,v := range a {
		fmt.Println("value",v)
	}
	
}

/*
nikhil@nikhil:~/job/100DaysOfCode/language/day1$ go run  for_range.go 
5
index 0
index 1
index 2
index 3
index 4
index 0 value 1
index 1 value 2
index 2 value 3
index 3 value 4
index 4 value 5
value 1
value 2
value 3
value 4
*/