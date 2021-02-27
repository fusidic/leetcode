package stack

// https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/solution/

// CQueue ...
type CQueue struct {
	p []int
	q []int
}

// Constructor ...
func Constructor() CQueue {
	return CQueue{}
}

// AppendTail ...
func (cq *CQueue) AppendTail(value int) {
	cq.p = append(cq.p, value)
}

func popFrom(q *[]int) (v int) {
	v, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]
	return
}

// DeleteHead ...
func (cq *CQueue) DeleteHead() int {
	if len(cq.q) > 0 {
		return popFrom(&cq.q)
	}
	for i := len(cq.p) - 1; i >= 0; i-- {
		cq.q = append(cq.q, cq.p[i])
	}
	cq.p = []int{}
	if len(cq.q) > 0 {
		return popFrom(&cq.q)
	}
	return -1
}
