package day9

import (
	"fmt"

	"github.com/mikeconroy/advent-of-code-24/utils"
)

func Run() (string, string) {
	input := utils.ReadFileIntoSlice("day9/input")
	return part1(input), part2(input)
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoublyLinkedList) append(node *Node) {
	if list.head == nil {
		list.head = node
	}
	if list.tail != nil {
		list.tail.next = node
		node.prev = list.tail
	}
	list.tail = node
}

func (list *DoublyLinkedList) print() {
	node := list.head
	for node.next != nil {
		fmt.Println("Node ID:", node.id, "Size:", node.size)
		node = node.next
	}
}

type Node struct {
	id   int
	size int
	next *Node
	prev *Node
}

func generateDll(diskMap string) *DoublyLinkedList {
	list := &DoublyLinkedList{}
	id := 0
	for block := range diskMap {
		size := int(diskMap[block] - '0')
		node := &Node{
			id:   id,
			size: size,
		}
		id += size
		list.append(node)
	}
	return list
}

func part1(input []string) string {
	list := generateDll(input[0])
	list.print()
	return fmt.Sprint(1)
}

func part2(input []string) string {
	return fmt.Sprint(2)
}
