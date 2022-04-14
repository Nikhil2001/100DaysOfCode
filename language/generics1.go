package main

import (
	"fmt"
	"reflect"
)

type Numeric interface {
	int | float32 | float64 | complex128 | complex64
}

func Print[T any](a T) {
	fmt.Println(a)
}

func add[T Numeric](a, b T) T {
	return a + b
}

func PrintAny(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	default:
		fmt.Println("Unhandled type")
	}
}

type GenericSlice[T any] []T

func (s GenericSlice[T]) Print() {
	fmt.Printf("type %T: ", s)
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

type Node[T any] struct {
	data T
	next *Node[T]
}

type list[T any] struct {
	head *Node[T]
}

func (l list[T]) add(data T) {

	if l.head == nil {
		l.head = &Node[T]{data, nil}
		return
	}

	if l.head.next == nil {
		l.head.next = &Node[T]{data, nil}
		return
	}

	temp := l.head
	l.head = l.head.next
	l.add(data)
	l.head = temp

}

func PrintSlice[T any](slice []T) {
	fmt.Printf("Printing the slice of %T \n", slice)
	for _, v := range slice {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func PrintSliceUsingReflect(v interface{}) {
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Slice {
		return
	}

	for i := 0; i < val.Len(); i++ {
		fmt.Print(val.Index(i).Interface(), " ")
	}

	fmt.Println()
}

func main() {
	Print(1)
	Print("nikhil")
	Print(true)
	Print(1.4 + 3i)
	Print(1.4)

	fmt.Println(add(1, 2), add(1.4+3i, 2i), add(1.4, 2.4))

	PrintAny(1)
	PrintAny("nikhil")
	PrintAny(true)
	PrintAny(1.4 + 3i)
	PrintAny(1.4)

	var gSlice1 = GenericSlice[int]{1, 2, 4}
	var gSlice2 = GenericSlice[float64]{1.0, 2.0, 4.1}
	fmt.Println(gSlice1, gSlice2)
	gSlice1.Print()
	gSlice2.Print()

	var nextNode = Node[int]{data: 2}
	var headNode = Node[int]{data: 1, next: &nextNode}
	var ll = list[int]{&headNode}
	ll.add(4)
	fmt.Println(ll.head, ll.head.next, ll.head.next.next)

	PrintSlice([]int{1, 2, 5})
	PrintSlice([]bool{true, false, true})

	PrintSliceUsingReflect([]int{1, 2, 5})
	PrintSliceUsingReflect([]bool{true, false, true})

}

/*
nikhil@nikhil:~/code$ go run main.go
1
nikhil
true
(1.4+3i)
1.4
3 (1.4+5i) 3.8
1
nikhil
Unhandled type
Unhandled type
Unhandled type
[1 2 4] [1 2 4.1]
type main.GenericSlice[int]: 1 2 4
type main.GenericSlice[float64]: 1 2 4.1
&{1 0xc000010260} &{2 0xc000010280} &{4 <nil>}
Printing the slice of []int
1 2 5
Printing the slice of []bool
true false true
1 2 5
true false true
*/

