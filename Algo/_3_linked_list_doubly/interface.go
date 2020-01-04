package _3_linked_list_doubly

type list interface {
	// 元素的数量
	size() int
	// 是否为空
	Empty() bool
	// 是否包含某个元素
	contains(element interface{}) bool
	// 返回index位置对应的元素
	get(index int) interface{}
	// 设置index位置的元素
	set(index int, element interface{}) interface{}
	// 往index位置添加元素
	add(index int, element interface{})
	// 删除index位置对应的元素
	remove(index int) interface{}
	// 查看元素的位置
	indexOf(element interface{}) int
	// 清除所有元素
	clear()
}
