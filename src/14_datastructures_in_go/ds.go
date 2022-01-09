package main

import (
	singlelinkedlist "14_datastructures_in_go/datastructures"
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

func main() {
	TestSingleLinkedList()
}
