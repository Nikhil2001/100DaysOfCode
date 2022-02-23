package main

import "fmt"

func sum(a,b int) int {
	return a+b
} 

func addAndSub(a,b int) (add int,sub int) {
	add = a+b
	sub = a-b
	return 
} 

func addSubMul(a,b int) (int, int, int) {
	return a+b,a-b,a*b
} 


func main(){
	fmt.Println(sum(1,2))
	fmt.Println(addAndSub(1,2))	
	fmt.Println(addSubMul(2,4))
}

