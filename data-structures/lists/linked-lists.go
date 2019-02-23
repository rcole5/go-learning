package lists

import (
	"errors"
)

type Node struct {
	Value interface{}
	NextNode *Node
	PrevNode *Node
}

type LinkedList struct {
	Head *Node
}

func New(init interface{}) *LinkedList {
	return &LinkedList{
		Head: &Node{
			Value:    init,
			NextNode: nil,
			PrevNode: nil,
		},
	}
}

func (node *Node) remove() {
	node.PrevNode.NextNode = node.NextNode
	node.NextNode.PrevNode = node.PrevNode
}

func tail(node *Node) *Node {
	if node.NextNode != nil {
		return tail(node.NextNode)
	}

	return node
}

func (list *LinkedList) Tail() *Node {
	return tail(list.Head)
}

func (list *LinkedList) Push(value interface{}) {
	node := tail(list.Head)

	node.NextNode = &Node {
		Value: value,
		NextNode: nil,
		PrevNode: node,
	}
}

func (list *LinkedList) Pop() Node {
	head := *list.Head

	list.Head = list.Head.NextNode
	list.Head.PrevNode = nil

	return head
}

func (list *LinkedList) Get(index int) (*Node, error) {
	if index == 0 {
		return list.Head, nil
	}

	current := list.Head.NextNode
	for x := 1; x <= index;  x++ {
		if x == index {
			return current, nil
		}
		current = current.NextNode
		if current == nil {
			break
		}
	}

	return nil, errors.New("element not found")
}

func (list *LinkedList) Delete(index int) error {
	element, err := list.Get(index)
	if err != nil {
		return err
	}

	element.remove()
	return nil
}
