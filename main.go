package main

import "github.com/rohitjb/heapProblems/maxheap"

func main() {
	items := []int{1, 23, 12, 9, 30, 2, 50}
	k := 6
	maxheap.KthSmallestElements(items, k)
}