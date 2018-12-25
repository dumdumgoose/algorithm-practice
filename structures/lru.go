package structures

type DoubleDirectionalNode struct {
	Prev *DoubleDirectionalNode
	Next *DoubleDirectionalNode
	Data interface{}
	Key  interface{}
}

type LRU struct {
	cache      map[interface{}]*DoubleDirectionalNode
	root, tail *DoubleDirectionalNode
	cap        int
}

func NewLRU(size int) *LRU {
	if size <= 0 {
		return nil
	}
	return &LRU{
		cache: make(map[interface{}]*DoubleDirectionalNode),
		cap:   size,
	}
}

func (l *LRU) Put(key interface{}, value interface{}) {
	newNode := &DoubleDirectionalNode{Data: value, Key: key}
	if l.Size() >= l.cap {
		l.evictLeastRecent()
	}
	if l.root == nil {
		l.root = newNode
		l.tail = l.root
	} else {
		l.root.Prev = newNode
		newNode.Next = l.root
		l.root = newNode
	}
	l.cache[key] = newNode
}

func (l *LRU) Get(key interface{}) (interface{}, bool) {
	if v, ok := l.cache[key]; ok {
		l.moveToHead(v)
		return v.Data, true
	}
	return nil, false
}

func (l *LRU) Size() int {
	return len(l.cache)
}

func (l *LRU) moveToHead(n *DoubleDirectionalNode) {
	if n.Prev != nil { // n is not root, so we need to move it to the head of the linked list
		// link n's prev and next
		if n.Next != nil { // check whether n is tail
			n.Next.Prev = n.Prev
		}
		n.Prev.Next = n.Next
		// move n to head
		n.Next = l.root
		n.Prev = nil
		l.root.Prev = n
		l.root = n
	}
}

// remove tail
func (l *LRU) evictLeastRecent() {
	if l.root == nil || l.tail == nil {
		return
	}

	// remove tail item from cache map
	delete(l.cache, l.tail.Key)

	// remove tail from linked list
	if l.tail.Prev != nil { // tail is not head
		temp := l.tail.Prev
		temp.Next = nil
		l.tail.Prev = nil
		l.tail = temp
	} else { // tail is head
		l.root = nil
		l.tail = nil
	}
}
