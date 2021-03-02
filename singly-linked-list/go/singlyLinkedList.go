package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

// Retturns pointers to the initialized memory for root node
var head = new(Node)

// Add the node
func (node *Node) addNode(value int) {
	if node == nil {
		head = node.addInitialNode(value)
	} else {
		node.appendNode(value)
	}
}

// add the first node
func (node *Node) addInitialNode(value int) *Node {
	addedNode := Node{Value: value, Next: nil}
	node = &addedNode
	return node
}

// Add the node at the begining of the list
func (node *Node) pushNode(value int) {
	var newNode Node
	newNode.Value = value
	newNode.Next = *&head
	head = &newNode
}

// Add the node at the end of the list
func (node *Node) appendNode(value int) {
	isNextNodePresent := true
	for isNextNodePresent == true {
		if node.Next != nil {
			node = node.Next
		} else {
			newNode := Node{Value: value, Next: nil}
			node.Next = &newNode
			isNextNodePresent = false
		}
	}
}

// Travel to each and every node and print its value
func (node *Node) traverse() {
	if node == nil {
		fmt.Println("Empty List!")
	}

	for node != nil {
		fmt.Print("->", *node)
		node = node.Next
	}
	fmt.Println()
}

// Search if the item is present in the list.
func (node *Node) searchItem(value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	}
	if node.Next == nil {
		return false
	}
	return node.Next.searchItem(value)
}

// Print the sile of the list
func (node *Node) listLength() int {
	if node == nil {
		return 0
	}
	size := 1
	for node.Next != nil {
		size++
		node = node.Next
	}
	return size
}

// Adds the node after the given position
func (node *Node) addInMiddle(value int, position int) {
	currentPos := 1
	for node.Next != nil {
		if position == currentPos {
			var newNode Node
			newNode.Value = value
			newNode.Next = node.Next
			node.Next = &newNode
		}
		currentPos++
		node = node.Next
	}
}

// Delete node At any given position
func (node *Node) deleteNode(position int) {

	// if node is empty
	if node == nil {
		fmt.Println("Node Empty. Cannot Delete")
	} else if position > node.listLength() { // if the given position is greater than the list size, no element can be deleted
		fmt.Println("Position more than list size.")
	} else if position == 1 { // if the node to delete is first node or the head node
		head = node.Next
	} else if position == node.listLength() { // if the node to delete is the last node i.e. position is the size of list
		for node != nil {
			if node.Next.Next == nil {
				node.Next = nil
			}
			node = node.Next
		}
	} else { // if the position of the node to delete is in between first and last node in the list
		preNode := node
		count := 1
		for node != nil {
			if count == position {
				node = preNode
				node.Next = node.Next.Next
				break
			}
			count++
			preNode = node
			node = node.Next
		}
	}
}

func main() {
	head = nil

	head.addNode(1)

	head.addNode(2)

	head.pushNode(5)

	head.appendNode(6)

	head.addInMiddle(9, 3)

	head.traverse()

	head.deleteNode(3)

	head.traverse()
}
