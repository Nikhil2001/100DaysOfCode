package main

import "fmt"

func Do(steps ...string) {
	defer fmt.Println("Done!")

	for _, v := range steps {
		defer fmt.Println(v)
	}

	fmt.Println("Starting....")
}

func Do2(steps ...string) {
	defer fmt.Println("Done!")

	defer func() {
		for _, v := range steps {
			fmt.Println(v)
		}
	}()

	fmt.Println("Starting....")
}

func main() {
	Do("see result", "Run Program", "Compile Code", "Save code", "Write Code", "open IDE")
	Do2("open IDE", "Write Code", "Save code", "Compile Code", "Run Program","see result")
}
