package utils

import "errors"

type Stack[T any] struct {
	elems []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		elems: make([]T, 0),
	}
}

func (stack *Stack[T]) Push(elem T) *Stack[T] {
	stack.elems = append(stack.elems, elem)
	return stack
}

func (stack *Stack[T]) Pop() (*Stack[T], error) {
	if len(stack.elems) == 0 {
		return stack, errors.New("no elements in stack")
	}

	stack.elems = stack.elems[:len(stack.elems)-1]
	return stack, nil
}

func (stack *Stack[T]) Top() (T, error) {
	if len(stack.elems) == 0 {
		return *new(T), errors.New("no elements in stack")
	}
	return stack.elems[len(stack.elems)-1], nil
}
