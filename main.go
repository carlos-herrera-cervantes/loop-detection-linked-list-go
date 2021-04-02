package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

// Returns the length of linked list
func (list linkedList) getLength() int {
	return list.length
}

// Prints the elements of linked list
func (list linkedList) printList() {
	current := list.head

	fmt.Printf("%v -> ", current.data)

	for current.next != list.head {
		current = current.next
		fmt.Printf("%v -> ", current.data)
	}

	fmt.Println()
}

// Adds new node in the linked list
func (list *linkedList) pushBack(node *node) {
	if list.head == nil {
		list.head = node
		list.tail = node
		node.next = list.head

		list.length++
		return
	}

	list.tail.next = node
	list.tail = node
	list.tail.next = list.head
	list.length++
}

// Finds the beginning node
func (list *linkedList) findBeginning() *node {
	slow := list.head
	fast := list.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			break
		}
	}

	if fast == nil || fast.next == nil {
		return nil
	}

	slow = list.head

	for slow != fast {
		slow = slow.next
		fast = fast.next
	}

	return fast
}

func main() {
	node1 := &node{data: 1}
	node2 := &node{data: 3}
	node3 := &node{data: 10}
	node4 := &node{data: 12}
	node5 := &node{data: 18}
	linkedList1 := linkedList{}

	linkedList1.pushBack(node1)
	linkedList1.pushBack(node2)
	linkedList1.pushBack(node3)
	linkedList1.pushBack(node4)
	linkedList1.pushBack(node5)

	result := linkedList1.findBeginning()

	if result != nil {
		fmt.Println("The beginning node is: ", result.data)
	} else {
		fmt.Println("There's no beginning node")
	}
}
