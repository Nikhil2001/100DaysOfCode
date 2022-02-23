package main

import "fmt"

func main(){
	for i:=0;i<5;i++ {
		 i2 := func() int {
			 return i*i
		 }()
		 fmt.Println(i,"square is",i2)
	}
}