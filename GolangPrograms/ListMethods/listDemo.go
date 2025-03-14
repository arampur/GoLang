package main

import "container/list"

func main() {
	var inlist list.List

	inlist.PushBack(1)
	inlist.PushBack(2)
	inlist.PushBack(3)
	inlist.PushFront(0)
	inlist.InsertAfter(4, inlist.Front())

	for element := inlist.Front(); element != nil; element = element.Next() {
		println(element.Value.(int))
	}
}
