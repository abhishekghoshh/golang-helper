package collections

import "fmt"

/*
https://gobyexample.com/generics
We can define methods on generic types just like we do on regular types,
but we have to keep the type parameters in place. The type is List[T], not List.
*/
func LinkedLists() {
	// manually creating the linked list object
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())

	// here type inferring will not work
	// we have to give the type manually to the function
	var strngs = NewList[string]()
	strngs.Push("Abhishek Ghosh")
	strngs.Push("Nasim Molla")
	strngs.Push("Bishal Mukherjee")
	strngs.Push("Abhishek Pal")
	fmt.Println(strngs.GetAll())
}

func NewList[T any]() List[T] {
	return List[T]{}
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
}

type Node[T any] struct {
	val  T
	next *Node[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &Node[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &Node[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}
