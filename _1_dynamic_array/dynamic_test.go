package dynamicarray

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var dynamicArray = NewDynamicArray(10)

func TestDynamicArray_Add(t *testing.T) {
	dynamicArray.Add(0, 1)
	dynamicArray.Print()
}

func TestDynamicArray_AddFirst(t *testing.T) {
	t1 := time.Now()
	goroutineNum := 100_000
	var wg sync.WaitGroup

	for i := 0; i < goroutineNum; i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			dynamicArray.AddFirst(ii)
		}(i)
	}
	wg.Wait()
	fmt.Println(dynamicArray.GetSize(), time.Since(t1).Seconds())
	//dynamicArray.Print()
}

func TestDynamicArray_AddTail(t *testing.T) {
	dynamicArray.AddTail(1)
	dynamicArray.AddTail(2)
	dynamicArray.Print()
}

func TestDynamicArray_Clear(t *testing.T) {
	dynamicArray.Clear()
	dynamicArray.Print()
}

func TestDynamicArray_Contains(t *testing.T) {
	for i := 0; i < 10; i++ {
		dynamicArray.AddTail(i)
	}

	fmt.Println(dynamicArray.Contains(73))
	fmt.Println(dynamicArray.Contains(7))
	dynamicArray.Print()
}

func TestDynamicArray_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		dynamicArray.AddFirst(i)
	}

	fmt.Println(dynamicArray.Get(3))
	fmt.Println(dynamicArray.Get(7))
	dynamicArray.Print()
}

func TestDynamicArray_GetCapacity(t *testing.T) {
	dynamicArray.Print()
	for i := 0; i < 12; i++ {
		dynamicArray.AddFirst(i)
	}
	dynamicArray.Print()
}

func TestDynamicArray_GetSize(t *testing.T) {
	dynamicArray.Print()
	for i := 0; i < 12; i++ {
		dynamicArray.AddFirst(i)
	}
	dynamicArray.Print()
}

func TestDynamicArray_IsEmpty(t *testing.T) {
	fmt.Println(dynamicArray.IsEmpty())
	for i := 0; i < 12; i++ {
		dynamicArray.AddFirst(i)
	}
	dynamicArray.Print()
	fmt.Println(dynamicArray.IsEmpty())
}

func TestDynamicArray_Remove(t *testing.T) {
	for i := 0; i < 12; i++ {
		dynamicArray.AddFirst(i)
	}
	dynamicArray.Print()
	dynamicArray.Remove(2)
	dynamicArray.Print()
}

func TestDynamicArray_Set(t *testing.T) {
	for i := 0; i < 12; i++ {
		dynamicArray.AddFirst(i)
	}
	dynamicArray.Print()
	fmt.Println(dynamicArray.Set(2, 22))
	dynamicArray.Print()
}

func TestDynamicArray_IndexOf(t *testing.T) {
	for i := 0; i < 12; i++ {
		dynamicArray.AddFirst(i)
	}
	dynamicArray.Print()
	fmt.Println(dynamicArray.IndexOf(2))
	dynamicArray.Remove(2)
	fmt.Println(dynamicArray.IndexOf(2))
}
