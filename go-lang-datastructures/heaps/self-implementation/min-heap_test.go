package minheap

import (
	"fmt"
	"testing"
)

func TestInsertMinHeap(T *testing.T) {
	minHeap := NewMinHeap()
	minHeap.Insert(20)
	minHeap.Insert(50)
	minHeap.Insert(-2)

	if res := minHeap.RemoveMin(); res != -2 {
		T.Errorf("Expected %v Got %v", -2, res)
	}

	if res := minHeap.RemoveMin(); res != 20 {
		T.Errorf("Expected %v Got %v", 20, res)
	}

	if res := minHeap.RemoveMin(); res != 50 {
		T.Errorf("Expected %v Got %v", 50, res)
	}

	if res := minHeap.RemoveMin(); res != -1 {
		T.Errorf("Expected %v Got %v", -1, res)
	}

	arr := make([]int, 2) //default zeros
	fmt.Printf("\nLeng of array : %d, capacity %d\n", len(arr), cap(arr))
	//fmt.Printf("\nLeng of array : %d, capacity %d", len(arr2), cap(arr2))

	arr = append(arr, 200)
	arr[1] = 20
	if arr[1] != 20 {
		T.Errorf("Expected arr[1]=%v, but got %v", 20, arr[1])
	}

	fmt.Printf("\nLeng of array : %d, capacity %d\n", len(arr), cap(arr))
	arr = arr[:len(arr)-1] //slice off the last element, so include everything except the last index
	fmt.Println(arr)

	val := minHeap.RemoveMin()
	fmt.Println("Val", val)

	fmt.Printf("value of the heap %+v", minHeap)

	val = minHeap.RemoveMin()
	fmt.Println("Val", val)

	fmt.Printf("value of the heap %+v", minHeap)

	val = minHeap.RemoveMin()
	fmt.Println("Val", val)

	fmt.Printf("value of the heap %+v", minHeap)

	val = minHeap.RemoveMin()
	fmt.Println("Val", val)

	fmt.Printf("value of the heap %+v", minHeap)
}
