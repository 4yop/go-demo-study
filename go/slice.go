package main

import "fmt"

func main () {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = i;
	}
	fmt.Println(arr)

	fmt.Println(arr[0:10])

	fmt.Println(len(arr))



	fmt.Println(cap(arr))


	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("len: ", len(slice))
	fmt.Println("cap: ", cap(slice))
	//改变切片长度
	slice = append(slice, 6)
	fmt.Println("after append operation: ")
	fmt.Println("len: ", len(slice))
	fmt.Println("cap: ", cap(slice)) //注意，底层数组容量不够时，会重新分配数组空间，通常为两倍
}

