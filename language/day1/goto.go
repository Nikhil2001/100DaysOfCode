package main

import "fmt"

func main() {

	var a string
	i := 1

start:
	for {

		switch {
		case i%2 == 0 && len(a) < 10:
			i++
			goto A
		case i%2 == 1 && len(a) < 10:
			i++
			goto B
		case len(a) > 10:
			break start
		}

	A:
		a += "a"
		continue
	B:
		a += "b"
		continue
	}
	fmt.Println(a)

}

/*
nikhil@nikhil:~/job/100DaysOfCode/language/day1$ go run goto.go
bababababaa
*/
