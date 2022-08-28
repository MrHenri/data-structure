package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	lenght int
}

// add node to the head of the list
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.lenght++
}

// add node to the tail of the list
func (l *linkedList) append(n *node) {
	if l.head == nil {
		l.head = n
		l.lenght++
		return
	}

	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	lastNode.next = n
	l.lenght++
}

func (l linkedList) printList() {
	if l.lenght == 0 {
		fmt.Print("Empty linked list")
	}

	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (l *linkedList) pop(index int) int {
	if index < 0 {
		fmt.Println("index has to be positive")
		return -1
	}
	if index >= l.lenght {
		fmt.Println("index out of range")
		return -1
	}

	currentNode := l.head
	if l.lenght == 1 {
		l.head = nil
		l.lenght--
		return currentNode.data
	}

	prevCurrent := currentNode
	for i := 0; i < index; i++ {
		prevCurrent = currentNode
		currentNode = currentNode.next
	}

	if currentNode.next == nil {
		prevCurrent.next = nil
		l.lenght--
		return currentNode.data
	}

	prevCurrent.next = currentNode.next
	l.lenght--
	return currentNode.data
}

func (l *linkedList) popByValue(value int) int {
	if l.lenght == 0 {
		fmt.Println("Empty linked list")
		return -1
	}

	currentNode := l.head
	if l.lenght == 1 {
		l.head = nil
		l.lenght--
		return currentNode.data
	}

	prevCurrent := currentNode
	for currentNode != nil {
		if currentNode.data == value {
			prevCurrent.next = currentNode.next
			l.lenght--
			return currentNode.data
		}
		prevCurrent = currentNode
		currentNode = currentNode.next
	}

	fmt.Println("Value is not in the list")
	return -1
}

func main() {
	myList := linkedList{}

	node1 := &node{data: 48}
	// node2 := &node{data: 18}
	// node3 := &node{data: 16}
	// node4 := &node{data: 4}
	// node5 := &node{data: 123}
	// node6 := &node{data: 76}

	myList.prepend(node1)
	// myList.prepend(node2)
	// myList.prepend(node3)
	// myList.prepend(node4)
	// myList.prepend(node5)
	// myList.prepend(node6)

	myList.printList()
	// fmt.Println(myList.pop(0))
	fmt.Println(myList.popByValue(48))
	// fmt.Println(myList.deleteNode(14))
	myList.printList()
}
