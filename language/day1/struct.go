package main

import "fmt"

type T20Player struct {
	name    string
	matches int
	runs    int
	wickets int
}

var cricketer1 = T20Player{name: "Venky Iyer", matches: 4, runs: 329, wickets: 7}
var cricketer2 = T20Player{name: "Virat kohli", matches: 97, runs: 3296}
var cricketer3 = T20Player{ matches: 9,name: "Ishan Kishan", runs: 296}

func main() {
	fmt.Println(cricketer1)
	fmt.Printf("%+v \n",cricketer2)
	fmt.Printf("%v \n",cricketer3)
}
