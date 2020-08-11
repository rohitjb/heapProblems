package main

import (
	"fmt"
	"github.com/rohitjb/heapProblems/maxheap"
	"github.com/rohitjb/heapProblems/minheap"
)

func main() {
}

func closetNumber() {
	item := []int{5, 6, 7, 8, 9}
	k := 3
	x := 7

}

func sortKsortedArray() {
	items := []int{6, 5, 3, 2, 8, 10, 9}
	k := 3
	fmt.Println(minheap.SortKSortedElement(items, k))
}

func printKthSmallestElements() {
	items := []int{1, 23, 12, 9, 30, 2, 50}
	k := 3
	maxheap.KthSmallestElements(items, k)
}