package main

import linkedlist "14_datastructures_in_go/datastructures"

func TestLinkedList() {
	list := &linkedlist.List{}
	list.Init()
	list.Insert(1)
	list.Insert(2)
	list.Print()
}

func main() {
	TestLinkedList()
}
