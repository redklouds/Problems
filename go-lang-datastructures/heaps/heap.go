package heap

/*
O(1) look up for Min and Max heap

O(logN) insert

golang container/heap
Interface has 3 methods

type interface interface {
	sort.Interfaceconst
	Push(x interface{})
	Pop() interface
}

*/

type Item struct {
	value interface{}
	pri   int
	index int
}

type MinHeap []*Item

// Less - determine which is more priority than another
func (pq MinHeap) Less(i, j int) bool {
	return pq[i].pri < pq[j].pri
}

// Swap - implementation of swap for the heap interface
func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push - implementation of push for the heap interface
func (pq *MinHeap) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop - implementation of pop for heap interface
func (pq *MinHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
