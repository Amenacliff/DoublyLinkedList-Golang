package main

import "log"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func (d *DoublyLinkedList) InitializeList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (d *DoublyLinkedList) AddNodeToFront(data int) {
	newNode := &Node{
		data: data,
	}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}

	d.length++
}

func (d *DoublyLinkedList) AddNodeToBack(data int) {
	previousData := d.tail
	newData := &Node{
		data: data,
	}
	if d.head == nil {
		d.head = newData
		d.tail = newData
	} else {
		newData.prev = d.tail
		previousData.next = newData
		d.tail = newData
	}

}

func (d *DoublyLinkedList) DisplayAllItemsInList() {
	currentNode := d.head
	for currentNode != nil {
		log.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {
	var linkedList = DoublyLinkedList{}
	linkedList.InitializeList()
	linkedList.AddNodeToFront(5)
	linkedList.AddNodeToFront(7)
	linkedList.AddNodeToFront(7)
	linkedList.AddNodeToFront(6)
	linkedList.AddNodeToFront(2)
	linkedList.AddNodeToFront(1)
	linkedList.AddNodeToBack(0)
	linkedList.AddNodeToBack(9)
	linkedList.DisplayAllItemsInList()
}
