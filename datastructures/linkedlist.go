package main

type Node struct {
	data int
	next *Node
}

func (n *Node) Add(data int) {
	n.data = data
	n.next = nil
}

func (n *Node) Get() int {
	return n.data
}

func (n *Node) Set(data int) {
	n.data = data
}

func (n *Node) Remove() {
	n.data = 0
	n.next = nil
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, tail: nil}
}

func (l *LinkedList) Add(data int) {
	newNode := &Node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

// Get returns the data of the first node in the list
func (l *LinkedList) Get() int {
	if l.head == nil {
		return 0
	}
	return l.head.Get()
}
