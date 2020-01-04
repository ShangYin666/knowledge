package dynamicarray

// IArray 数组接口定义
type IArray interface {
	// 基本方法
	GetSize() int     // 元素的数量
	IsEmpty() bool    // 是否为空
	GetCapacity() int // 数组容量

	// 添加元素
	AddFirst(element interface{})       // 在数组首位插入元素
	AddTail(element interface{})        // 在数组末尾插入元素
	Add(index int, element interface{}) // index位置添加元素

	// 删除元素
	Remove(index int) interface{} // 删除index位置对应的元素
	Clear()                       // 清除所有元素

	// 修改
	Set(index int, element interface{}) interface{} // 设置index位置的元素

	// 查询方法
	Contains(element interface{}) bool // 是否包含某个元素
	Get(index int) interface{}         // 返回index位置对应的元素
	IndexOf(element interface{}) int   // 查找元素的位置
	Print()
}
