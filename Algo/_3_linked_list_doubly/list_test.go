package _3_linked_list_doubly

import (
	"fmt"
	"testing"
)

var l = New()

func Test_doublyLinkedList_Add(t *testing.T) {

	for i := 0; i < 10; i++ {
		l.AddHead(i)
	}
	l.Traverse()
	fmt.Println(l.IsEmpty(), l.GetSize())
}
