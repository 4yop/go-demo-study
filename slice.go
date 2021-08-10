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

	slice = append(slice,arr... )

	fmt.Println("after append operation: ")
	fmt.Println("len: ", len(slice))
	fmt.Println("cap: ", cap(slice)) //注意，底层数组容量不够时，会重新分配数组空间，通常为两倍

	var arr1 [5]int = [5]int{1,2,3,4,5}

	fmt.Println("arr1的全部",arr1)
	fmt.Println("arr1的index=0~3",arr1[:3])
	fmt.Println("arr1的index=2~结尾",arr1[2:])
	fmt.Println("arr1的index=2~4",arr1[2:4])
	fmt.Println("arr1的index=0~结尾",arr1[:])

	//s1 := []int{1,2,3,4}
	//s2 := make([]int,0)
}

