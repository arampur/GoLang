package main

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// head -> 10 -> 18 -> 17 -> nil

func (list *LinkedList) insertAtFront(data int) {
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
	}

}

func (list *LinkedList) insertAtBack(data int) {

}

func main() {

}
