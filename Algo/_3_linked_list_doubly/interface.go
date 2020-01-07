package _3_linked_list_doubly

type IList interface {
	// normal action
	GetSize() int
	IsEmpty() bool
	Traverse()

	// Add node
	AddHead(element interface{})        // 头节点插入元素
	AddTail(element interface{})        // 尾节点插入元素
	Add(index int, element interface{}) // 将元素插入到指定的位置

	// delete node
	Remove(index int) interface{} // 移除某个位置的元素
	Clear()                       // 清空链表

	// update node
	Set(index int, element interface{}) interface{} // 修改指定位置的元素

	// select node
	Container(element interface{}) bool // 是否包含某个元素
	IndexOf(element interface{}) int    // 查看某个元素的位置
}
