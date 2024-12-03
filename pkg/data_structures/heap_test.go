package data_structures

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
)

func TestInsert(t *testing.T) {
	heap := Heap[int]{
		arr: []int{},
		compare: func(a, b int) int {
			if a < b {
				return -1
			} else if a > b {
				return 1
			} else {
				return 0
			}
		},
		head: 0,
		tail: 0,
	}

	heap.Insert(9)
	heap.Insert(4)
	heap.Insert(5)
	heap.Insert(11)
	if heap.arr[0] != 11 {
		t.Errorf("Insert 11 where max is 9. Want 11, got %d.\nHeap: %v", heap.arr[0], heap.arr)
	}

	for i := range 3 {
		heap.Insert(i)
	}
	if heap.arr[0] != 11 {
		t.Errorf("Insert 3 thru 0 where max is 11. Want 11, got %d.\nHeap: %v", heap.arr[0], heap.arr)
	}

	heap.Insert(64).
		Insert(13).
		Insert(10)

	if heap.arr[0] != 64 {
		t.Errorf("Insert 64, 13, 9  where max is 11. Want 64, got %d.\nHeap: %v", heap.arr[0], heap.arr)
	}
}

func TestMultipleInsert(t *testing.T) {
	minHeap := func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	heap := NewHeap(minHeap)

	heap.Insert(9).
		Insert(4).
		Insert(5).
		Insert(11).
		Insert(64).
		Insert(13).
		Insert(10)

	expectedVals := []int{64, 13, 11, 10, 9, 5}
	for _, expected := range expectedVals {
		if val, _ := heap.Pop(); val != expected {
			t.Errorf("Expected %d got %d.\nHeap %v", expected, val, heap.arr)
		}
	}
}

func TestDump(t *testing.T) {
	minHeap := func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	heap := NewHeap(minHeap)
	heap.Insert(9).
		Insert(4).
		Insert(5).
		Insert(11).
		Insert(64).
		Insert(13).
		Insert(10)

	sorted, err := heap.Dump()
	if err != nil {
		t.Errorf("%v", err)
	}

	if len(sorted) != 7 {
		t.Errorf("Expected a sorted array of 7 elements, got %d elements.\n", len(sorted))
	}

	if !heap.IsEmpty() {
		t.Errorf("Expected empty heap after dump!")
	}

	heap.Insert(9).
		Insert(4).
		Insert(5).
		Insert(11).
		Insert(64).
		Insert(13).
		Insert(10)

	for i := 0; !heap.IsEmpty(); i++ {
		val, err := heap.Pop()
		if err != nil {
			t.Errorf("%v", err)
		}

		if val != sorted[i] {
			t.Errorf("Expected %v Got %v\n", val, sorted[i])
		}
	}
}

func TestDuplicates(t *testing.T) {
	minHeap := func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	heap := NewHeap(minHeap)

	heap.Insert(4).
		Insert(4).
		Insert(4).
		Insert(4).
		Insert(4).
		Insert(3).
		Insert(0)

	heap.Dump()
}

func TestSimpleDiff(t *testing.T) {
	minHeap := func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	heap1 := NewHeap(minHeap)
	heap2 := NewHeap(minHeap)

	heap1.Insert(1).
		Insert(2).
		Insert(3).
		Insert(4).
		Insert(5).
		Insert(6).
		Insert(7)

	heap2.Insert(0).
		Insert(1).
		Insert(2).
		Insert(3).
		Insert(4).
		Insert(5).
		Insert(6)

	a, _ := heap1.Dump()
	b, _ := heap2.Dump()

	var sum int64
	for i := range a {
		diff := a[i] - b[i]
		if diff < 0 {
			diff = 0 - diff
		}
		sum += int64(diff)
	}
	if sum != 7 {
		t.Errorf("Expected 7, got %d", sum)
	}

	heap1.Insert(1).
		Insert(2).
		Insert(3).
		Insert(4).
		Insert(5).
		Insert(6).
		Insert(7)

	heap2.Insert(6).
		Insert(5).
		Insert(4).
		Insert(3).
		Insert(2).
		Insert(1).
		Insert(0)

	a, _ = heap1.Dump()
	b, _ = heap2.Dump()

	sum = 0
	for i := range a {
		diff := a[i] - b[i]
		if diff < 0 {
			diff = 0 - diff
		}
		sum += int64(diff)
	}
	if sum != 7 {
		t.Errorf("Expected 7, got %d", sum)
	}

}

func genRandMinHeap(size int) (*Heap[*big.Int], error) {
	minHeap := func(a, b *big.Int) int {
		return b.Cmp(a)
	}
	heap := NewHeap(minHeap)

	for range size {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(size*3)))
		if err != nil {
			return nil, fmt.Errorf("Expected to be able to generate random big int.\n%w\n", err)
		}

		heap.Insert(n)
	}
	return heap, nil
}

func checkOrder(results []*big.Int) (bool, error) {
	for i := range results {
		if i+2 == len(results) {
			break
		}
		if results[i].Cmp(results[i+1]) > 0 {
			return false, fmt.Errorf("Expected %v to be less than %v.\nHeap: %v\n", results[i], results[i+1], results)
		}
	}
	return true, nil
}

func TestDump10(t *testing.T) {
	heap, err := genRandMinHeap(10)
	if err != nil {
		t.Errorf("Expected to be able to genrate random heap.\n%v", err)
	}
	results, err := heap.Dump()
	if err != nil {
		t.Errorf("Expected to be able to Dump() heap. %v\n", err)
	}

	ok, err := checkOrder(results)
	if !ok {
		t.Errorf("%v", err)
	}
}

func TestDump100(t *testing.T) {
	heap, err := genRandMinHeap(100)
	if err != nil {
		t.Errorf("Expected to be able to genrate random heap.\n%v", err)
	}
	results, err := heap.Dump()
	if err != nil {
		t.Errorf("Expected to be able to Dump() heap. %v\n", err)
	}

	ok, err := checkOrder(results)
	if !ok {
		t.Errorf("%v", err)
	}
}

func TestDump100_000(t *testing.T) {
	heap, err := genRandMinHeap(100_000)
	if err != nil {
		t.Errorf("Expected to be able to genrate random heap.\n%v", err)
	}
	results, err := heap.Dump()
	if err != nil {
		t.Errorf("Expected to be able to Dump() heap. %v\n", err)
	}

	ok, err := checkOrder(results)
	if !ok {
		t.Errorf("%v", err)
	}
}
