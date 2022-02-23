package main

import "fmt"

func doubleValue(a *int) {
	*a = 2 * (*a)

}

func double(a int) {
	a = 2 * a
}

func main() {
	a := 3
	double(a)
	fmt.Println(a) // no change 
	doubleValue(&a)
	fmt.Println(a)

}
