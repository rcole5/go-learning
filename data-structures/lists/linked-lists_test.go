package lists

import (
	"testing"
)

func TestNew(t *testing.T) {
	list := New("Hello")
	if list.Head.Value != "Hello" {
		t.Fail()
	}

	if list.Head.NextNode != nil {
		t.Fail()
	}
}

func TestPush(t *testing.T) {
	list := New("Hello")

	if list.Head.Value != "Hello" {
		t.Fatal("Head value != Hello")
	}

	if list.Head.NextNode != nil {
		t.Fatal("Head next not empty")
	}

	list.Push("World")
	node := list.Tail()

	if node.Value != "World" {
		t.Fatal("Value != world")
	}

	if node.NextNode != nil {
		t.Fatal("Next node not nil")
	}

	if node.PrevNode != list.Head {
		t.Fatal("Prev node not head")
	}
}

func TestGet(t *testing.T) {
	list := New("Hello")
	list.Push("World")
	list.Push("!")

	element, err := list.Get(2)
	if err != nil {
		t.Fatal(err)
	}

	if element.Value != "!" {
		t.Fatal("Element not !")
	}

	_, err = list.Get(3)
	if err == nil {
		t.Fatal("element should be out of bounds")
	}
}

func TestPop(t *testing.T) {
	list := New("Hello")
	list.Push("World")
	list.Push("!")

	oldHead := list.Pop()
	if oldHead.Value != "Hello" {
		t.Fatal("incorrect old head")
	}

	if list.Head.Value != "World" {
		t.Fatal("incorrect new head")
	}

	if list.Head.PrevNode != nil {
		t.Fatal("new head prev has not been removed")
	}
}

func TestDelete(t *testing.T) {
	list := New("Hello")
	list.Push("World")
	list.Push("!")

	err := list.Delete(1)
	if err != nil {
		t.Fatal(err)
	}

	element, err := list.Get(1)
	if err != nil {
		t.Fatal(err)
	}

	if element.Value != "!" {
		t.Fatal("expected !")
	}
}
