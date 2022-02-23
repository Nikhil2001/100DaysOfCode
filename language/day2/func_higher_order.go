package main

import "fmt"

func ChangeNums(nums []int,f func(int) int) func() {
	for i,v := range nums{
		nums[i]= f(v)
	}
	return func(){
		fmt.Println(nums)
	}
}

var f2 = func(a int) int{
	return a*a
} 

var f3 = func(a int) int{
	return a*a*a
}

func main(){

	nums := []int{1,2,3}
	f := ChangeNums(nums,f2)
	f()

	nums = []int{4,5,6}
	g := ChangeNums(nums,f3)
	g()

}