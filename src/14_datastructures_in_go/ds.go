package main

import (
	doublelinkedlist "14_datastructures_in_go/doublelinkedlist"
	singlelinkedlist "14_datastructures_in_go/singlelinkedlist"
	"fmt"
)

func TestSingleLinkedList() {
	list := &singlelinkedlist.List{}
	list.Init()
	list.Insert(1)
	list.Insert(2)
	list.Print()
	list.Delete(1)
	list.Print()
	list.Insert(22)
	list.Insert(222)
	fmt.Println(list.Search(222))
	fmt.Println(list.Search(223))
	list.Print()
}

func TestDoubleLinkedList() {
	list := &doublelinkedlist.List{}
	list.Init()
	list.Insert(1)
	list.Insert(2)
	list.Print()
	list.PrintReverse()
	list.Delete(1)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)
	list.Print()
	list.PrintReverse()
	fmt.Println(list.RevSearch(5))
	fmt.Println(list.Search(6))
	list.Delete(4)
	list.Delete(6)
	list.PrintReverse()
}

func main() {
	fmt.Println("Testing Single LinkedList")
	TestSingleLinkedList()
	fmt.Println("Testing Double LinkedList")
	TestDoubleLinkedList()
}
