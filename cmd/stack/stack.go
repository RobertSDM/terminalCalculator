package stack

import "fmt"

// A representation of the concept LIFO (Last In First Out)
type Stack[T comparable] struct {
	list []T
	Len  int
}

// Append an element to the stack
func (s *Stack[T]) Add(str T) {
	s.list = append(s.list, str)
	s.Len++
}

// Removes the top element of the stack
//
// If the stacks length is 0 raise panic.
// To avoid it use [Stack.HashLen]
func (s *Stack[T]) Pop() T {
	if s.HasLen() {
		value := s.Top()
		s.list = s.list[:s.Len-1]
		s.Len--
		return value
	}

	panic("cannot pop the slice, not enough elements")
}

// Returns the element at the top of the stack
//
// If there is no element a panic is raised.
// To avoid it use the [Stack.HasLen] method
func (s *Stack[T]) Top() T {
	if s.HasLen() {
		return s.list[len(s.list)-1]
	}

	panic("no element at the top of the stack")
}

// Verify if the stack has some element in it
//
// It's recomended to use this method before [Stack.Pop] and [Stack.Top]
// to avoid panics
func (s *Stack[T]) HasLen() bool {
	return s.Len > 0
}

func (s Stack[T]) String() string {
	return fmt.Sprint(s.list)
}

// Create a stack
func CreateStack[U comparable]() *Stack[U] {
	return &Stack[U]{
		list: make([]U, 0),
	}
}
