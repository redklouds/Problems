package minheap

type minHeap struct {
	heap []int
	size int
}

func NewMinHeap() minHeap {
	return minHeap{
		heap: make([]int, 2),
		size: 0,
	}
}

//heap insert is O(LogN)
func (h *minHeap) Insert(data int) {
	if h.size == 0 {
		//h.heap = make([]int, 2)
		h.heap[1] = data
		h.size++
	} else {
		h.heap = append(h.heap, data) //insert at the last index
		h.size++
		//insert at the tail
		//shift to keep the min heap rule
		h.shiftUp(h.size, h.size/2) // i/2 will always get the parent
	}

}

func (h *minHeap) shiftUp(childIdx int, parentIdx int) {
	if h.heap[childIdx] < h.heap[parentIdx] && parentIdx != 0 {
		//need to swap
		temp := h.heap[parentIdx]
		h.heap[parentIdx] = h.heap[childIdx]
		h.heap[childIdx] = temp
		h.shiftUp(parentIdx, parentIdx/2)
	}
}

func (h *minHeap) RemoveMin() int {
	var retMinVal int
	if h.size == 0 {
		//only the first index
		//and cap(h.heap) == 2? reserved 0th index for clarity
		//defined with two
		retMinVal = -1
	} else {
		//remove the top/root at index 1 decrement and shift down
		retMinVal = h.heap[1]
		h.heap[1] = h.heap[h.size]
		h.heap = h.heap[:h.size]
		h.size--

		h.shiftDown(1)
	}

	return retMinVal
}

func (h *minHeap) shiftDown(parentIdx int) {
	if h.size  >1 {

		leftN := parentIdx * 2
		rightN := parentIdx*2 + 1

		if h.heap[parentIdx] >= h.heap[leftN] && h.heap[parentIdx] >= h.heap[rightN] {
			if h.heap[leftN] < h.heap[rightN] {
				//go left
				temp := h.heap[leftN]
				h.heap[leftN] = h.heap[parentIdx]
				h.heap[parentIdx] = temp
				parentIdx = leftN
			} else {
				temp := h.heap[rightN]
				h.heap[rightN] = h.heap[parentIdx]
				h.heap[parentIdx] = temp
				parentIdx = rightN
			}
			h.shiftDown(parentIdx)
		}
	}

}
