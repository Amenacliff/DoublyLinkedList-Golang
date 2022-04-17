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

	d.length++

}

func (d *DoublyLinkedList) DisplayAllItemsInListForward() {
	currentNode := d.head
	for currentNode != nil {
		log.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (d *DoublyLinkedList) DisplayAllItemsInListBackward() {
	currentNode := d.tail
	for currentNode != nil {
		log.Println(currentNode.data)
		currentNode = currentNode.prev
	}
}

func (d *DoublyLinkedList) DeleteNode(index int) {
	if d.length == 0 {
		return
	}

	currentNode := d.head

	counter := 0

	for counter != index {
		currentNode = currentNode.next
		counter++
	}

	log.Println(currentNode.next, d.tail)
	if currentNode != d.tail {
		previousNode := currentNode.prev
		previousNode.next = currentNode.next
	} else {
		previousNode := currentNode.prev
		previousNode.next = nil
		d.tail = previousNode
	}

	d.length--
}

func (d *DoublyLinkedList) AddItemToListByIndex(index, data int) {
	if d.length == 0 || index == 0 {
		d.AddNodeToFront(data)
	} else {
		halfLengthOfList := d.length / 2

		if index < halfLengthOfList {
			counter := 0
			currentNode := d.head

			for counter != index {
				currentNode = currentNode.next
				counter++
			}

			previousNode := currentNode.prev
			newNode := &Node{
				data: data,
			}

			newNode.prev = previousNode
			newNode.next = currentNode
			previousNode.next = newNode
			currentNode.prev = newNode

		} else {
			counter := 0
			currentNode := d.tail

			for counter != index {
				currentNode = currentNode.prev
				counter++
			}

			previousNode := currentNode.prev
			newNode := &Node{
				data: data,
			}

			newNode.prev = previousNode
			newNode.next = currentNode
			previousNode.next = newNode
			currentNode.prev = newNode
		}
		d.length++
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
	linkedList.AddItemToListByIndex(1, 8)
	linkedList.AddNodeToBack(0)
	linkedList.AddNodeToBack(9)
	linkedList.DeleteNode(3)
	linkedList.DisplayAllItemsInListForward()
	log.Println(".....................")
	linkedList.DisplayAllItemsInListBackward()
}
