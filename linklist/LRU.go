package linklist

// LRUCache ...
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkNode
	head, rear *DLinkNode
}

// DLinkNode ...
type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}

func initDLinkNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	head := initDLinkNode(0, 0)
	rear := initDLinkNode(0, 0)
	head.next = rear
	rear.prev = head
	return LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*DLinkNode{},
		head:     head,
		rear:     rear,
	}
}

// Get ...
func (lc *LRUCache) Get(key int) int {
	if _, ok := lc.cache[key]; !ok {
		return -1
	}
	lc.float(lc.cache[key])

	return lc.cache[key].value
}

// Put ...
func (lc *LRUCache) Put(key int, value int) {
	if node, ok := lc.cache[key]; !ok {
		newNode := initDLinkNode(key, value)
		lc.cache[key] = newNode
		lc.headInsert(newNode)
		if lc.size > lc.capacity {
			deleteNode := lc.tailRemove()
			delete(lc.cache, deleteNode.key)
		}
	} else {
		node.value = value
		lc.float(node)
	}
}

func (lc *LRUCache) float(node *DLinkNode) {
	lc.removeNode(node)
	lc.headInsert(node)
}

func (lc *LRUCache) headInsert(node *DLinkNode) {
	node.next = lc.head.next
	node.next.prev = node
	node.prev = lc.head
	lc.head.next = node
	lc.size++
}

func (lc *LRUCache) removeNode(removed *DLinkNode) {
	removed.next.prev = removed.prev
	removed.prev.next = removed.next
	lc.size--
}

func (lc *LRUCache) tailRemove() *DLinkNode {
	removed := lc.rear.prev
	lc.removeNode(removed)
	return removed
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
