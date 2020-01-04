package dynamicarray

import (
	"fmt"
	"strings"
)

// DynamicArray 动态数组
type DynamicArray struct {
	len  int
	data []interface{}
}

const (
	// DefaultCapacity 默认数组容量
	DefaultCapacity = 10
	// ElementNotFound 查询不到该元素
	ElementNotFound = -1
)

// NewDynamicArray 初始化动态数组容量
func NewDynamicArray(capacity uint) IArray {
	if capacity < DefaultCapacity {
		capacity = DefaultCapacity
	}
	return &DynamicArray{
		data: make([]interface{}, capacity),
	}
}

// GetSize 获取数组容量
func (d *DynamicArray) GetSize() int {
	return d.len
}

// IsEmpty 判断数组是否为空
func (d *DynamicArray) IsEmpty() bool {
	return d.len == 0
}

// GetCapacity 获取数组容量
func (d *DynamicArray) GetCapacity() int {
	return len(d.data)
}

// AddFirst 在数组首位添加元素
func (d *DynamicArray) AddFirst(element interface{}) {
	d.Add(0, element)
}

// AddTail 在数组末尾添加元素
func (d *DynamicArray) AddTail(element interface{}) {
	d.Add(d.len, element)
}

// Add 数组中添加元素
func (d *DynamicArray) Add(index int, element interface{}) {
	d.rangeCheckForAdd(index)
	// 当数组的长度和容量相等时候需要考虑扩容问题
	if d.GetCapacity() == d.len {
		// 扩容
		d.expansionCapacity()
	}

	/**
	当插入第一个元素  1
	d.len ==0    index == 0

	当插入第二个元素  2
	d.len == 1    index==0     d.data[0]=1  移位   d.data[0]=0   d.data[1]=1   再将新元素写入d.data[index]的位置

	当插入第三个元素 3
	d.len = 2     index == 0   d.data[0]=2  d.data[1]=1  移位  d.data[0]=0  d.data[1]=2 d.data[2]=1  再将新元素写入d.data[index]的位置
	*/
	for i := d.len - 1; i >= index; i-- {
		d.data[i+1] = d.data[i]
	}
	d.data[index] = element
	d.len++
}

// Remove 移除数组中指定索引的元素
func (d *DynamicArray) Remove(index int) interface{} {
	d.rangeCheck(index)
	oldValue := d.data[index]
	/**
	d.data[0]=1  d.data[1]=2  d.data[2]=3
	如果删除index=0的元素，则将index 为1和2的元素向前移动一位覆盖前一个元素
	*/
	for i := index; i < d.len-1; i++ {
		d.data[i] = d.data[i+1]
	}
	d.len--
	//  如果数组的长度是容量的一半，则进行缩容
	if d.len == d.GetCapacity()>>1 {
		d.trimCapacity()
	}
	return oldValue
}

// Clear 清空数组
func (d *DynamicArray) Clear() {
	d.len = 0
}

// Set 修改数组中指定位置的元素
func (d *DynamicArray) Set(index int, element interface{}) interface{} {
	d.rangeCheck(index)
	oldVal := d.data[index]
	d.data[index] = element
	return oldVal
}

// Contains 判断数组中是否包含某个元素
func (d *DynamicArray) Contains(element interface{}) bool {
	return d.IndexOf(element) != ElementNotFound
}

// Get 根据数组索引获取元素
func (d *DynamicArray) Get(index int) interface{} {
	d.rangeCheck(index)
	return d.data[index]
}

// indexOf 根据元素查找数组索引
func (d *DynamicArray) IndexOf(element interface{}) int {
	if d.IsEmpty() {
		return ElementNotFound
	}
	for index := 0; index < d.len; index++ {
		if d.data[index] == element {
			return index
		}
	}
	return ElementNotFound
}
func (d DynamicArray) outOfBounds() {
	panic("out of range")
}
func (d *DynamicArray) rangeCheck(index int) {
	if index < 0 || index >= d.len {
		d.outOfBounds()
	}
}
func (d *DynamicArray) rangeCheckForAdd(index int) {
	if index < 0 || index > d.len {
		d.outOfBounds()
	}
}

// expansionCapacity 数组扩容
func (d *DynamicArray) expansionCapacity() {
	// 老数组
	capacity := d.GetCapacity() << 1
	newArray := make([]interface{}, capacity)
	for i := 0; i < d.len; i++ {
		newArray[i] = d.data[i]
	}
	d.data = newArray
}

// trimCapacity 数组缩容
func (d *DynamicArray) trimCapacity() {
	// 老数组
	capacity := d.GetCapacity() >> 1
	newArray := make([]interface{}, capacity)
	for i := 0; i < d.len; i++ {
		newArray[i] = d.data[i]
	}
	d.data = newArray
}

func (d *DynamicArray) Print() {
	str := fmt.Sprintf("容量为：%d，元素数量为：%d，数组为：[ ", d.GetCapacity(), d.len)
	for i := 0; i < d.len; i++ {
		str += fmt.Sprintf("%v->%v , ", i, d.data[i])
	}
	trimStr := strings.TrimRight(str, " ,")
	trimStr += "]"
	fmt.Println(trimStr)
}
