package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		if i == 3 { 
			continue // skips the rest of loop from here
		}
		fmt.Println(i)
	}

loop:
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(j, " ")
			if j == 2 {
				fmt.Println()
				continue loop //skips printing after 2
			}
		}
		fmt.Println()
	}

	for i := 0; i < 10; i++ {
	loop2:
		for j := 0; j < i; j++ {
			if j == 2 {
				continue loop2 // will not print on j == 2 iteration.
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}


/* OUTPUT :
nikhil@nikhil:~/job/100DaysOfCode/language/day1$ go  run continue.go
0
1
2
4

0 
0 1 
0 1 2 
0 1 2 
0 1 2 
0 1 2 
0 1 2 
0 1 2 
0 1 2 

0 
0 1 
0 1 
0 1 3 
0 1 3 4 
0 1 3 4 5 
0 1 3 4 5 6 
0 1 3 4 5 6 7 
0 1 3 4 5 6 7 8 

*/