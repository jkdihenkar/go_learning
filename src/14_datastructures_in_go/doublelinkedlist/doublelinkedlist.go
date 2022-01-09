package doublelinkedlist

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Init() {
	l.head = nil
	l.tail = nil
}

func (l *List) Insert(key interface{}) {
	node := &Node{nil, nil, key}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}
}

func (l *List) Delete(key interface{}) {
	if l.head == nil {
		return
	}
	if l.head.key == key {
		l.head = l.head.next
		return
	}
	curr := l.head
	for curr.next != nil {
		if curr.next.key == key {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

func (l *List) Print() {
	if l.head == nil {
		return
	}
	curr := l.head
	for curr != nil {
		fmt.Printf("%v ", curr.key)
		curr = curr.next
	}
	fmt.Println()
}

func (l *List) PrintReverse() {
	if l.tail == nil {
		return
	}
	curr := l.tail
	for curr != nil {
		fmt.Printf("%v ", curr.key)
		curr = curr.prev
	}
	fmt.Println()
}

func (l *List) Search(key interface{}) bool {
	if l.head == nil {
		return false
	}
	curr := l.head
	for curr != nil {
		if curr.key == key {
			return true
		}
		curr = curr.next
	}
	return false
}

func (l *List) RevSearch(key interface{}) bool {
	if l.tail == nil {
		return false
	}
	curr := l.tail
	for curr != nil {
		if curr.key == key {
			return true
		}
		curr = curr.prev
	}
	return false
}
