package main

import "errors"

type PriorityQueue[T interface{}] struct {
	elems map[int32][]T
}

func (queue *PriorityQueue[T]) AddElement(priority int32, elem T) {
	value, ok := queue.elems[priority]
	if !ok {
		queue.elems[priority] = make([]T, 1)
		queue.elems[priority][0] = elem
		return
	}

	value = append(value, elem)
	queue.elems[priority] = value
}

func (queue *PriorityQueue[T]) SortedKeys() []int32 {
	keys := make([]int32, 0, len(queue.elems))
	for k := range queue.elems {
		keys = append(keys, k)
	}
	return keys
}

func (queue *PriorityQueue[T]) GetMaxPriority() (int32, error) {
	if len(queue.elems) == 0 {
		return 0, errors.New("No elements in queue")
	}

	return queue.SortedKeys()[0], nil
}

func (queue *PriorityQueue[T]) Pop() (priority int32, elem T, err error) {
	priority, err = queue.GetMaxPriority()
	if err != nil {
		return
	}

	elem = queue.elems[priority][0]
	if len(queue.elems[priority]) == 1 {
		delete(queue.elems, priority)
	}
	queue.elems[priority] = queue.elems[priority][1:]
	return
}
