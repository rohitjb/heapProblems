package closestNumberHeap
import (
	"container/heap"
	"fmt"
)

type ClosetNumber struct {
	value    int
	difference int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*ClosetNumber

// An IntHeap is a min-heap of ints.

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i].difference < h[j].difference }
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ClosetNumber))
}

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func PrintMinHeap(list []int, k int) {
	items := IntHeap(list)
	heap.Init(&items)
	for items.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(&items))
	}
}
