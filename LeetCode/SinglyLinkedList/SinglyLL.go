package main

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// head -> 10 -> 18 -> 17 -> nil

func (l *LinkedList) addToHead(data int) {
	node := &Node{data: data}
	node.next = l.head
	l.head = node
}

func (l *LinkedList) addToEnd(data int) {
	node := &Node{data: data}
	if l.head == nil {
		l.head = node
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (l *LinkedList) lastNode() *Node {
	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current
}

func (l *LinkedList) addAfterNode(prevNode *Node, data int) {
	if prevNode == nil {
		return
	}
	node := &Node{data: data}
	node.next = prevNode.next
	prevNode.next = node
}

func (l *LinkedList) iterateList() {
	current := l.head
	for current != nil {
		println(current.data)
		current = current.next
	}
}

func main() {
	ll := &LinkedList{}
	ll.addToEnd(10)
	ll.addToEnd(18)
	ll.addToEnd(17)
	ll.addAfterNode(ll.lastNode(), 20)
	ll.iterateList()
}
