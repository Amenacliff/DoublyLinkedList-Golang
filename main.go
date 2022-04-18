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

func (d *DoublyLinkedList) AddItemToListByIndex(index, data int) {
	if d.length <= index {
		log.Println("Index Is Greater than Length of the List, hence it does not exist")
		return
	}

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
			counter := d.length - 1
			currentNode := d.tail

			for counter != index {
				currentNode = currentNode.prev
				counter--
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

func (d *DoublyLinkedList) DeleteItemFromListByIndex(index int) {
	if d.length == 0 {
		return
	}

	if d.length < index {
		log.Panicf("Index Is Greater than Lenght of the List, hence it does not exist")
		return
	}

	if d.length == 1 {
		d.head = nil
		d.tail = nil
		d.length = 0
		return
	}

	halfOfTheList := d.length / 2

	if index < halfOfTheList {
		counter := 0
		currentNode := d.head

		for counter != index {
			currentNode = currentNode.next
			counter++
		}

		previousNode := currentNode.prev
		nextNode := currentNode.next
		nextNode.prev = previousNode
		previousNode.next = currentNode.next
		currentNode.prev = previousNode
	} else {
		counter := d.length - 1
		currentNode := d.tail

		for counter != index {
			currentNode = currentNode.prev
			counter--
		}

		previousNode := currentNode.prev
		nextNode := currentNode.next
		nextNode.prev = previousNode
		previousNode.next = currentNode.next
		currentNode.prev = previousNode
	}

	d.length--

}

func (d *DoublyLinkedList) DeleteLastNode() {
	lastNode := d.tail

	if lastNode == nil {
		return
	}

	if lastNode == d.head {
		d.head = nil
		d.tail = nil
		d.length = 0
	}

	previousNode := lastNode.prev

	previousNode.next = nil

	d.tail = previousNode

	d.length--
}

func (d *DoublyLinkedList) ReverseList() {
	temp := &Node{}
	tail := d.head

	currentNode := d.head

	for currentNode != nil {
		temp = currentNode.prev
		currentNode.prev = currentNode.next
		currentNode.next = temp
		currentNode = currentNode.prev
	}
	if temp != nil {
		d.head = temp.prev
		d.tail = tail
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
	linkedList.AddItemToListByIndex(7, 8)
	linkedList.DisplayAllItemsInListForward()
	log.Println(".....................")
	linkedList.ReverseList()
	linkedList.DisplayAllItemsInListBackward()
}
