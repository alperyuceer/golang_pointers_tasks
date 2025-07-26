package odev4

import (
	"fmt"
	"go_pointers_tasks/odev1"
)

func Odev4_1() {
	odev1.Space_between_tasks("Odev4_1")

	var head *Node

	nodeEkle(&head, 10)
	nodeEkle(&head, 20)
	nodeEkle(&head, 30)

	listeyiYazdir(head)

}

type Node struct {
	data int
	next *Node
}

func nodeOlustur(data int) *Node {
	yeniNode := new(Node)
	yeniNode.data = data
	return yeniNode

}

func nodeEkle(head **Node, data int) {
	yeniNode := nodeOlustur(data)

	if *head == nil {
		*head = yeniNode
	} else {
		current := *head
		for current.next != nil {
			current = current.next
		}
		current.next = yeniNode
	}

}

func listeyiYazdir(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()

}
