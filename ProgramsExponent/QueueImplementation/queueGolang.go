package main

import "fmt"

type Queue []int

func main() {
	var q Queue
	q.isEmpty()

	q.enQueue(1)
	q.enQueue(2)
	q.enQueue(4)
	q.enQueue(14)

	q.deQueue()
	q.deQueue()

	q.enQueue(3)

	q.printElements(q)
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) enQueue(item int) {
	*q = append(*q, item)
}

func (q *Queue) deQueue() {
	if q.isEmpty() {
		fmt.Println("Queue is empty..")
	} else {
		index := len(*q)
		item := (*q)[0]
		fmt.Printf("Dequeued %d\n", item)
		*q = (*q)[1:index]
	}
}

func (q *Queue) printElements(queue []int) {
	fmt.Println("Queue elements after enqueue and dequeue")
	for i := 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
