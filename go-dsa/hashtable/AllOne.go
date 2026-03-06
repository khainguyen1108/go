package hashtable

type Node struct {
	freq int
	keys map[string]struct{}
	prev *Node
	next *Node
}

type AllOne struct {
	head *Node
	tail *Node
	mp   map[string]*Node
}

func Constructor() AllOne {
	head := &Node{freq: 0}

	tail := &Node{freq: 0}
	head.next = tail
	tail.prev = head

	return AllOne{
		head: head,
		tail: tail,
		mp:   map[string]*Node{},
	}
}

func (this *AllOne) Inc(key string) {
	if nodeCurr, ok := this.mp[key]; ok {
		freq := nodeCurr.freq
		delete(nodeCurr.keys, key)
		nextNode := nodeCurr.next
		if nextNode == this.tail || nextNode.freq != freq+1 {
			newNode := &Node{
				freq: freq + 1,
				keys: map[string]struct{}{},
			}

			newNode.keys[key] = struct{}{}
			newNode.prev = nodeCurr
			newNode.next = nextNode
			nodeCurr.next = newNode
			nextNode.prev = newNode

			this.mp[key] = newNode
		} else {
			nextNode.keys[key] = struct{}{}
			this.mp[key] = nextNode
		}

		if len(nodeCurr.keys) == 0 {
			this.removeNode(nodeCurr)
		}
	} else {
		firstNode := this.head.next

		if firstNode == this.tail || firstNode.freq > 1 {
			newNode := &Node{
				freq: 1,
				keys: map[string]struct{}{},
			}
			newNode.prev = this.head
			newNode.next = firstNode
			this.head.next = newNode
			firstNode.prev = newNode
			newNode.keys[key] = struct{}{}
			this.mp[key] = newNode
		} else {
			firstNode.keys[key] = struct{}{}
			this.mp[key] = firstNode
		}
	}
}

func (this *AllOne) Dec(key string) {

	if nodeCurr, ok := this.mp[key]; ok {
		if nodeCurr.freq == 1 {
			delete(this.mp, key)
		}

		freq := nodeCurr.freq
		delete(nodeCurr.keys, key)
		prevNode := nodeCurr.prev

		if prevNode == this.head || prevNode.freq != freq-1 {
			newNode := &Node{
				freq: freq - 1,
				keys: map[string]struct{}{},
			}

			newNode.keys[key] = struct{}{}
			newNode.prev = prevNode
			newNode.next = nodeCurr
			prevNode.next = newNode
			nodeCurr.prev = newNode
			this.mp[key] = newNode
		} else {
			prevNode.keys[key] = struct{}{}
			this.mp[key] = prevNode
		}

		if len(nodeCurr.keys) == 0 {
			this.removeNode(nodeCurr)
		}
	}
}

func (this *AllOne) GetMaxKey() string {

	if this.tail.prev == this.head {
		return ""
	}

	for k := range this.tail.prev.keys {
		return k
	}

	return ""
}

func (this *AllOne) GetMinKey() string {

	if this.head.next == this.tail {
		return ""
	}

	for k := range this.head.next.keys {
		return k
	}

	return ""
}

func (this *AllOne) removeNode(node *Node) {

	prevNode := node.prev
	nextNode := node.next

	prevNode.next = nextNode
	nextNode.prev = prevNode
}
