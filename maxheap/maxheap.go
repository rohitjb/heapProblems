package maxheap

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    int // The value of the item; arbitrary.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue)Top() int { return pq[pq.Len()-1].value }

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].value > pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int) {
	item.value = value
	heap.Fix(pq, item.index)
}

func PrintHeap(pq PriorityQueue) {
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%d ", item.value)
	}
}

func CreateMaxHeap(items []int) PriorityQueue {
	pq := make(PriorityQueue, len(items))
	for index, value := range items {
		pq[index] = &Item{
			value:    value,
			index:    index,
		}
	}
	heap.Init(&pq)
	return pq
}

func KthSmallestElements(list []int, k int) {
	queue := CreateMaxHeap([]int{})

	for index, value := range list {
		item := &Item{
			value: value,
			index: index,
		}
		heap.Push(&queue, item)
		if queue.Len() > k {
			_ = heap.Pop(&queue)
		}
	}

	for queue.Len() > 0 {
		val := heap.Pop(&queue).(*Item).value
		fmt.Println(val)
	}
}

func SortKsortedElementInDecendingOrder(list []int, k int) []int {
	queue := CreateMaxHeap([]int{})
	sortedArray := []int{}

	for index, value := range list {
		item := &Item{
			value: value,
			index: index,
		}
		heap.Push(&queue, item)
		if queue.Len() > k {
			element := heap.Pop(&queue).(*Item).value
			sortedArray = append(sortedArray, element)
		}
	}

	for queue.Len() > 0 {
		val := heap.Pop(&queue).(*Item).value
		sortedArray = append(sortedArray, val)
	}
	return sortedArray
}