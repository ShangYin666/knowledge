package _2_linked_list_single

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var l = New()

func Test_singleLinkedList(t *testing.T) {
	l.AddHead(3)
	l.AddHead(2)
	l.AddHead(1)

	l.AddTail(5)
	l.AddTail(6)
	l.AddTail(7)

	l.Add(3, 4)
	//l.Add(0, 4)
	oldVal := l.Set(3, 44)
	fmt.Println(oldVal)

	fmt.Println(l.Container(4))
	fmt.Println(l.IndexOf(44))

	fmt.Println(l.Remove(0))
	fmt.Println(l.Remove(4))
	fmt.Println(l.Remove(3))
	l.Traverse()
}

func TestSingleLinkedList_Bench(t *testing.T) {
	var wg sync.WaitGroup
	goroutineNum := 10000
	t1 := time.Now()
	for i := 0; i < goroutineNum; i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			l.AddHead(ii)
			l.Add(0, ii)
			l.AddTail(ii)
		}(i)
	}
	wg.Wait()
	fmt.Println(l.GetSize(), time.Since(t1).Seconds())
	l.Traverse()
	l.Remove(0)
	fmt.Println(l.IsEmpty())
	l.Clear()
}
