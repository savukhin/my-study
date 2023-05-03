package utils

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

func (stack *Stack[T]) Pop() *Stack[T] {
	stack.elems = stack.elems[:len(stack.elems)-1]
	return stack
}

func (stack *Stack[T]) Top() T {
	return stack.elems[len(stack.elems)-1]
}
