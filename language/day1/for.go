package main

import "fmt"

func main() {

	// for loop as while
	i := 0
	for i < 5 {
		fmt.Println("Iteration:", i)
		i++
	}

	//traditional For loop

	for i := 0; i < 5; i++ {
		fmt.Println("Traditional For Loop, Iteration:", i)
	}

	// infinite for loop but breaks on condition
	i = 0
	for {
		fmt.Println("For Loop Iteration:", i)
		if i == 5 {
			break
		}
		i++
	}

	for j := 0; ; j++ {
		fmt.Println("For Loop with no condition:", j)
		if j == 5 {
			break
		}
	}

	i = 0
	for ; i < 5; i++ {
		fmt.Println("For loop with no initialization,Iteration:", i)

	}

	for k := 0; k < 5; {
		fmt.Println("For loop with no last statement, iteration :", i)
		k += 2
	}

}


/*
nikhil@nikhil:~/job/100DaysOfCode/language/day1$ go run for.go 
Iteration: 0
Iteration: 1
Iteration: 2
Iteration: 3
Iteration: 4
Traditional Iteration: 0
Traditional Iteration: 1
Traditional Iteration: 2
Traditional Iteration: 3
Traditional Iteration: 4
For Loop Iteration: 0
For Loop Iteration: 1
For Loop Iteration: 2
For Loop Iteration: 3
For Loop Iteration: 4
For Loop Iteration: 5
For Loop with no condition: 0
For Loop with no condition: 1
For Loop with no condition: 2
For Loop with no condition: 3
For Loop with no condition: 4
For Loop with no condition: 5
For loop with no initialization,Iteration: 0
For loop with no initialization,Iteration: 1
For loop with no initialization,Iteration: 2
For loop with no initialization,Iteration: 3
For loop with no initialization,Iteration: 4
For loop with no last statement, iteration : 5
For loop with no last statement, iteration : 5
For loop with no last statement, iteration : 5
*/

