package main

import (
	"errors"
	"fmt"
)

type stack struct {
	arr []int
	top int
}

func New() stack {
	newStack := stack{[]int{}, -1}
	return newStack
}

func (s *stack) Push(n int) {
	s.arr = append(s.arr, n)
	s.top = s.top + 1
}

func (s *stack) Pop() (int, error) {
	if !s.IsEmpty() {
		last := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		s.top = s.top - 1
		return last, nil
	}
	return 0, errors.New("invalid pop, cannot pop on Empty stack!")
}

func (s *stack) Top() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("invalid Top, Empty stack!")
	}
	return s.arr[s.top], nil
}

func (s *stack) IsEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *stack) Size() int {
	if s.top == -1 {
		return 0
	}
	return s.top + 1
}

func (s *stack) Show() {
	for _, v := range s.arr {
		fmt.Print(v, " ")
	}
	defer fmt.Printf("\n")
}

func main() {
	var s stack = New()
	s.Show()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Show()
	fmt.Println(s.Size())
	fmt.Println(s.Top())
	s.Show()
	fmt.Println(s.Pop())
	s.Show()
	fmt.Println(s.IsEmpty())

	s.Pop()
	s.Pop()
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Top())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
}

/*
nikhil@nikhil:~/job/100DaysOfCode/stacks$ go run stack.go

1 2 3
3 <nil>
1 2 3
3 <nil>
1 2
false
true
-1 invalid Top, Empty stack!
0 invalid pop, cannot pop on Empty stack!
true

*/
