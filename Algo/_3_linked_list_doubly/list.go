package _3_linked_list_doubly

import (
	"fmt"
	"strings"
)

type node struct {
	data interface{}
	prev *node
	next *node
}

type doublyLinkedList struct {
	size int
	head *node
	last *node // 保存last节点，方便元素查找
}

const (
	// ElementNotFound 查询不到该元素
	ElementNotFound = -1
)

func New() IList {
	return &doublyLinkedList{
		head: nil,
		last: nil,
	}
}
func (d *doublyLinkedList) GetSize() int {
	return d.size
}

func (d *doublyLinkedList) IsEmpty() bool {
	return d.size == 0
}

func (d *doublyLinkedList) Traverse() {
	node := d.head
	fmt.Println("-------------------------------------------")
	i := 0
	str := ""
	for node != nil {
		fmt.Printf("size=%d ,index=%v , nodePtr=%p , nodeData=%v , head=%v , last=%v \n", d.GetSize(), i, node, node.data, d.head, d.last)
		str += fmt.Sprintf("%v->", node.data)
		node = node.next
		i++
	}
	fmt.Println()
	fmt.Printf("双链表元素格式为：%d,链表结构为：%v \n\n", d.size, strings.TrimRight(str, "->"))
	fmt.Println("-------------------------------------------")
}

func (d *doublyLinkedList) AddHead(element interface{}) {
	d.Add(0, element)
}

func (d *doublyLinkedList) AddTail(element interface{}) {
	d.Add(d.size, element)
}

func (d *doublyLinkedList) Add(index int, element interface{}) {
	d.rangeCheckAdd(index)
	newNode := &node{data: element}

	if index == d.size { // 追加插入  index == d.size 共有2种情况， 头节点为空，头节点不为空
		oldLast := d.last
		d.last = newNode
		newNode.prev = oldLast
		if oldLast == nil { // 空链表
			d.head = d.last
		} else { //
			oldLast.next = newNode
		}
	} else { // 往前面插入元素
		// 获取原来的index位置的元素
		next := d.getNode(index)
		// 存储旧元素的前驱节点，新元素作为该节点的后继元素
		prev := next.prev // nil
		fmt.Println(prev, newNode, next)
		// 将该元素作为新节点的后继节点，则该节点变为新节点的next
		next.prev = newNode

		// 将新节点和前驱后继节点链接起来
		newNode.next = next // 0
		newNode.prev = prev // nil

		// 只有一个节点
		if prev == nil {
			d.head = newNode
		} else {
			// 原节点上一个节点的next,即为新添加节点
			prev.next = newNode
		}
	}
	d.size++
}

func (d *doublyLinkedList) Remove(index int) interface{} {
	d.rangeCheck(index)
	node := d.getNode(index)
	oldValue := node.data

	prev := node.prev
	next := node.next

	// 头删除
	if prev != nil {
		prev.next = next
	} else {
		d.head = next
	}

	// 尾删除
	if next != nil {
		next.prev = prev
	} else {
		d.last = next
	}
	return oldValue
}

func (d *doublyLinkedList) Clear() {
	d.size = 0
	d.head = nil
	d.last = nil
}

func (d *doublyLinkedList) Set(index int, element interface{}) interface{} {
	d.rangeCheck(index)
	node := d.getNode(index)
	oldVal := node.data
	node.data = element
	return oldVal
}

func (d *doublyLinkedList) Container(element interface{}) bool {
	return d.IndexOf(element) != ElementNotFound
}

func (d *doublyLinkedList) IndexOf(element interface{}) int {
	node := d.head
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

func (d *doublyLinkedList) rangeCheck(index int) {
	if index < 0 || index >= d.size {
		panic("Out of range")
	}
}

func (d *doublyLinkedList) rangeCheckAdd(index int) {
	if index < 0 || index > d.size {
		panic("Out of range")
	}
}

func (d *doublyLinkedList) getNode(index int) *node {
	// 1. 全量循环查找
	//node := d.head
	//i := 0
	//for node != nil {
	//	if i == index {
	//		break
	//	}
	//	node = node.next
	//	i++
	//}
	//return node

	// 2. 将链表切割成两部分做判断，时间比第一种循环节省一般时间
	//判断当前的index在size的前半部分还是后半部分
	if index < d.size>>1 { // 在前半部分
		node := d.head
		// 这里使用 < ，因为index是从1开始的
		for i := 0; i < index; i++ {
			node = node.next
		}
		return node
	} else { // 在后半部分  从后往前循环
		node := d.last
		for i := d.size - 1; i > index; i-- {
			node = node.prev
		}
		return node
	}

}
