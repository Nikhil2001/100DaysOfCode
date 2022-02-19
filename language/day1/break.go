package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		if i == 3 { // break the loop
			break
		}
		fmt.Println(i)
	}

loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(j, " ")
			if j == 2 {
				fmt.Println()
				break loop // breaks outside "loop"
			}
		}
		fmt.Println()
	}

	for i := 0; i < 5; i++ {
	loop2:
		for j := 0; j < i; j++ {
			fmt.Print(j, " ")
			if j == 2 {
				break loop2 // breaks inside loop "loop2", same as just break here
			}
		}
		fmt.Println()
	}
}


/* OUTPUT :
nikhil@nikhil:~/job/100DaysOfCode/language/day1$ go run break.go 
0
1
2

0 
0 1 
0 1 2 

0 
0 1 
0 1 2 
0 1 2 

*/