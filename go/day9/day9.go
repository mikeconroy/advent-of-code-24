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
	for node != nil {
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
	for block := range diskMap {
		size := int(diskMap[block] - '0')
		var id int
		if block%2 == 0 {
			id += block / 2
		} else {
			id = -1
		}
		node := &Node{
			id:   id,
			size: size,
		}

		list.append(node)
	}
	return list
}

func rearrangeWholeFiles(list *DoublyLinkedList) {
	last := list.tail
	for last != nil {
		for last.id == -1 {
			last = last.prev
		}
		space := nextFreeBlock(list.head, last)

		for space != last {
			space = nextFreeBlock(space, last)
			if space.size >= last.size {
				break
			}
			space = space.next
		}
		if space.id == -1 && space != last {
			if space.size == last.size {
				space.id = last.id
				last.id = -1
			} else {
				newSpace := &Node{
					id:   -1,
					size: space.size - last.size,
					prev: space,
					next: space.next,
				}
				space.size = last.size
				space.id = last.id
				last.id = -1
				space.next = newSpace
				newSpace.next.prev = newSpace
			}
		}
		last = last.prev
	}
}

func rearrange(list *DoublyLinkedList) {
	last := list.tail
	space := nextFreeBlock(list.head, list.tail)
	for {
		if last.id == -1 {
			last = last.prev
		}
		if last.size > space.size {
			space.id = last.id
			last.size -= space.size
		} else if last.size < space.size {
			newSpace := &Node{id: -1, size: space.size - last.size, prev: space, next: space.next}
			space.next.prev = newSpace
			space.id = last.id
			space.size = last.size
			space.next = newSpace
			last.id = -1
			last = last.prev
		} else {
			space.id = last.id
			last.id = -1
			last = last.prev
		}
		space = nextFreeBlock(space, last)

		if space.id != -1 || space == last {
			break
		}
	}

}

func nextFreeBlock(from *Node, upTo *Node) *Node {
	node := from
	for node.id != -1 {
		if node == upTo {
			return node
		}
		node = node.next
	}
	return node
}

func calculateChecksum(list *DoublyLinkedList) int {
	checksum := 0
	pos := 0
	node := list.head
	for node != nil {
		if node.id != -1 {
			for i := pos; i < pos+node.size; i++ {
				checksum += (i * node.id)
			}
		}
		pos += node.size
		node = node.next
	}
	return checksum
}

func part1(input []string) string {
	list := generateDll(input[0])
	rearrange(list)
	result := calculateChecksum(list)
	return fmt.Sprint(result)
}

func part2(input []string) string {
	list := generateDll(input[0])
	rearrangeWholeFiles(list)
	result := calculateChecksum(list)
	return fmt.Sprint(result)
}
