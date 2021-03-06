package sort

import "container/heap"

// Heap 堆
type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push value into Heap
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop value out of Heap
func (h *Heap) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

// 小顶堆
type minHeap struct {
	Heap
}

func (h minHeap) Less(i, j int) bool {
	return h.Heap[i] < h.Heap[j]
}

// 大根堆
type maxHeap struct {
	Heap
}

func (h maxHeap) Less(i, j int) bool {
	return h.Heap[i] > h.Heap[j]
}

// MedianFinder ...
type MedianFinder struct {
	minHeap *minHeap
	maxHeap *maxHeap
}

// Constructor ...
func Constructor() MedianFinder {
	minHeap := &minHeap{}
	maxHeap := &maxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)

	return MedianFinder{
		minHeap: minHeap,
		maxHeap: maxHeap,
	}
}

// AddNum ...
func (mf *MedianFinder) AddNum(num int) {
	if mf.minHeap.Len() == mf.maxHeap.Len() {
		heap.Push(mf.maxHeap, num)
		heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
	} else {
		heap.Push(mf.minHeap, num)
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	}
}

// FindMedian ...
func (mf *MedianFinder) FindMedian() float64 {
	if mf.minHeap.Len() != mf.maxHeap.Len() {
		return float64(mf.minHeap.Heap[0])
	}
	return (float64(mf.minHeap.Heap[0]) + float64(mf.maxHeap.Heap[0])) / 2
}
