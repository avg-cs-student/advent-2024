package data_structures

import "fmt"

type Heap[T any] struct {
	arr     []T
	compare func(T, T) int
	head    int
	tail    int
}

func NewHeap[T any](f func(T, T) int) *Heap[T] {
	return &Heap[T]{
		arr:     []T{},
		compare: f,
		head:    0,
		tail:    0,
	}
}

func (h *Heap[T]) IsEmpty() bool {
	return h.head == h.tail
}

func (h *Heap[T]) Size() int {
	return h.tail
}

func (h *Heap[T]) Insert(elem T) *Heap[T] {
	if len(h.arr) >= h.tail+1 {
		h.arr[h.tail] = elem
	} else {
		h.arr = append(h.arr, elem)
	}
	h.tail++

	for i := h.tail - 1; h.compare(h.arr[i], h.arr[parent(i)]) > 0; {
		h.arr[parent(i)], h.arr[i] = h.arr[i], h.arr[parent(i)]
		i = parent(i)
	}

	return h
}

func (h *Heap[T]) Pop() (T, error) {
	if h.IsEmpty() {
		var noop T
		return noop, fmt.Errorf("Heap is empty!")
	}

	popped := h.arr[h.head]
	h.tail--

	if h.IsEmpty() {
		return popped, nil
	}

	h.arr[h.head] = h.arr[h.tail]

	for root := h.head; ; {
		hasRightChild := right(root) < h.tail
		hasLeftChild := left(root) < h.tail

		if hasRightChild && hasLeftChild {
			shouldSwap := h.compare(h.arr[root], h.arr[right(root)]) < 0 || h.compare(h.arr[root], h.arr[left(root)]) < 0
			if !shouldSwap {
				break
			}

			if h.compare(h.arr[right(root)], h.arr[left(root)]) > 0 {
				h.arr[root], h.arr[right(root)] = h.arr[right(root)], h.arr[root]
				root = right(root)
			} else if h.compare(h.arr[left(root)], h.arr[right(root)]) >= 0 {
				h.arr[root], h.arr[left(root)] = h.arr[left(root)], h.arr[root]
				root = left(root)
			}
			continue
		} else if hasLeftChild {
			shouldSwap := h.compare(h.arr[root], h.arr[left(root)]) < 0
			if !shouldSwap {
				break
			}

			h.arr[root], h.arr[left(root)] = h.arr[left(root)], h.arr[root]
			root = left(root)
			continue
		}

		break
	}

	return popped, nil
}

func (h *Heap[T]) Dump() ([]T, error) {
	sorted := []T{}
	for !h.IsEmpty() {
		val, err := h.Pop()
		if err != nil {
			return nil, err
		}
		sorted = append(sorted, val)
	}

	return sorted, nil
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}
