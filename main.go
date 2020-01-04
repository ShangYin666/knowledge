package main

import "fmt"

func main() {
	a := 1537
	fmt.Println(a << 1)
	fmt.Println(a + a/4<<1)

	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
}
