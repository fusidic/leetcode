package sort

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	bucket := make(map[int]int)
	for _, v := range nums {
		bucket[v]++
	}
	h := &mHeap{}
	heap.Init(h)
	// 堆排序
	for key, freq := range bucket {
		heap.Push(h, [2]int{key, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := []int{}
	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(h).([2]int)[0])
	}
	return ret
}

// h[i][0] 维护数值
// h[i][1] 表示该值的个数
type mHeap [][2]int

func (h mHeap) Len() int {
	return len(h)
}
func (h mHeap) Less(x, y int) bool {
	// 比较二者频次的大小
	return h[x][1] < h[y][1]
}
func (h mHeap) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}
func (h *mHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *mHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[0 : l-1]
	return x
}
