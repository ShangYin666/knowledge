package _2_linked_list_single

import (
	"fmt"
	"sync"
)

type node struct {
	data interface{}
	next *node
}
type singleLinkedList struct {
	size  int
	head  *node
	mutex sync.RWMutex
}

const (
	// ElementNotFound 查询不到该元素
	ElementNotFound = -1
)

// 初始化链表
func New() IList {
	return &singleLinkedList{
		size: 0,
		head: nil,
	}
}
func (s *singleLinkedList) GetSize() int {
	return s.size
}

func (s *singleLinkedList) IsEmpty() bool {
	return s.size == 0
}

func (s *singleLinkedList) Traverse() {
	node := s.head
	i := 0
	fmt.Println("------------------------------------------------")
	for node != nil {
		fmt.Printf("%v %v %p %v \n", i, node.data, node, node)
		node = node.next
		i++
	}
	fmt.Println("------------------------------------------------")
}

func (s *singleLinkedList) AddHead(element interface{}) {
	s.Add(0, element)
}

func (s *singleLinkedList) AddTail(element interface{}) {
	s.Add(s.size, element)
}

func (s *singleLinkedList) Add(index int, element interface{}) {
	s.rangeCheckAdd(index)
	newNode := &node{data: element}

	// 空链表
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.head == nil {
		s.head = newNode
	} else {
		if index == 0 {
			// 链表头插入元素
			newNode.next = s.head
			s.head = newNode
		} else {
			// 链表尾部或中部插入元素
			prevNode := s.getNode(index - 1)
			nextNode := prevNode.next
			prevNode.next = newNode
			newNode.next = nextNode
		}
	}
	s.size++
}

func (s *singleLinkedList) Remove(index int) interface{} {
	s.rangeCheck(index)
	node := s.head
	if index == 0 { // 删除头
		s.head = s.head.next
	} else {
		prevNode := s.getNode(index - 1)
		// 1 3 4
		// index == 1  prevnode == 1
		prevNode.next = prevNode.next.next
	}
	s.size--
	return node.data
}

func (s *singleLinkedList) Clear() {
	s.size = 0
	s.head = nil
}

func (s *singleLinkedList) Set(index int, element interface{}) interface{} {
	s.rangeCheck(index)
	node := s.getNode(index)
	oldData := node.data
	node.data = element
	return oldData
}

func (s *singleLinkedList) Container(element interface{}) bool {
	return s.IndexOf(element) != ElementNotFound
}

func (s *singleLinkedList) IndexOf(element interface{}) int {
	node := s.head
	i := 0
	for node.next != nil {
		if node.data == element {
			return i
		}
		node = node.next
		i++
	}
	return ElementNotFound
}

func (s *singleLinkedList) rangeCheck(index int) {
	if index < 0 || index >= s.size {
		panic("Out of range")
	}
}
func (s *singleLinkedList) rangeCheckAdd(index int) {
	if index < 0 || index > s.size {
		panic("Out of range")
	}
}

func (s *singleLinkedList) getNode(index int) *node {
	node := s.head
	i := 0
	for node.next != nil {
		if index == i {
			break
		}
		node = node.next
		i++
	}
	return node
}
