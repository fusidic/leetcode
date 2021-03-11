package quiz

import "container/heap"

/**
 * the medians
 * @param operations int整型二维数组 ops
 * @return double浮点型一维数组
 */
func flowmedian(operations [][]int) []float64 {
	// write code here
	ret := []float64{}
	mf := contructor()
	for _, operation := range operations {
		if operation[0] == 1 {
			mf.addNum(operation[1])
		} else if operation[0] == 2 {
			ret = append(ret, mf.getMedian())
		}
	}
	return ret
}

type Heap []int

// Implement heap method

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// 小顶堆
type minHeap struct {
	Heap
}

func (h minHeap) Less(i, j int) bool {
	return h.Heap[i] < h.Heap[j]
}

// 大顶堆
type maxHeap struct {
	Heap
}

func (h maxHeap) Less(i, j int) bool {
	return h.Heap[i] > h.Heap[j]
}

type medianFinder struct {
	minHeap *minHeap
	maxHeap *maxHeap
}

func contructor() medianFinder {
	minHeap := &minHeap{}
	maxHeap := &maxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)
	return medianFinder{
		minHeap: minHeap,
		maxHeap: maxHeap,
	}
}

func (mf *medianFinder) addNum(x int) {
	if mf.minHeap.Len() == mf.maxHeap.Len() {
		heap.Push(mf.maxHeap, x)
		heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
	} else {
		heap.Push(mf.minHeap, x)
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	}
}

func (mf *medianFinder) getMedian() float64 {
	if mf.maxHeap.Len() == mf.minHeap.Len() {
		return (float64(mf.minHeap.Heap[0]) + float64(mf.maxHeap.Heap[0])) / 2
	}
	// 防止堆大小为0
	if mf.maxHeap.Len() == 0 {
		return -1
	}
	return float64(mf.minHeap.Heap[0])
}
