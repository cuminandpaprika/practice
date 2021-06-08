package main

import (
	"errors"
	"fmt"
)

var IndexOutOfBoundsError = errors.New("index is greater than size of linked list")

type SingleNode struct {
	Value string
	Next  *SingleNode
}

type DoubleNode struct {
	Value string
	Last  *DoubleNode
	Next  *DoubleNode
}

// GetItem is O(N) time complexity
func (s *SingleNode) GetItem(index int) (string, error) {
	pointer := s
	for i := 0; i < index; i++ {
		if pointer.Next == nil {
			return "", IndexOutOfBoundsError
		}
		pointer = pointer.Next
	}
	return pointer.Value, nil
}

// DeleteItem is O(N) time complexity
func (s *SingleNode) DeleteItem(index int) error {
	pointer := s
	for i := 0; i < index-1; i++ {
		if pointer.Next == nil {
			return IndexOutOfBoundsError
		}
		pointer = pointer.Next
	}

	// We've found the item to operate on

	// Handle index out of bounds
	if pointer.Next == nil {
		return IndexOutOfBoundsError
	}
	pointer.Next = pointer.Next.Next
	return nil
}

// MakeSingleLinkedList is O(N) time complexity
func MakeSingleLinkedList(values []string) SingleNode {
	head := SingleNode{}
	pointer := &head
	for i := range values {
		pointer.Value = values[i]
		nextItem := SingleNode{}
		pointer.Next = &nextItem
		pointer = &nextItem
	}
	return head
}

func main() {
	values := []string{"a", "b", "c"}
	singleLinkedList := MakeSingleLinkedList(values)
	err := singleLinkedList.DeleteItem(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(singleLinkedList.GetItem(0))
	fmt.Println(singleLinkedList.GetItem(1))
	fmt.Println(singleLinkedList.GetItem(2))
}
