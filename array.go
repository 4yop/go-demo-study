package main

import (
	"fmt"
)

func main()  {
	var arr [5] int;
	fmt.Printf("%p\n",&arr)
	for i := 0 ;i < len(arr);i++ {
		arr[i] = i + 1
		fmt.Printf("%p\n",&arr[i])
	}
	fmt.Printf("%p\n",&arr)

	var arr1 [5] int = [5]int{1,2,3,4,5};
	var arr2 = [5]int{1,2,3,4,5}
	arr3 := [5]int{1,2,3,4,5}

	arr4 := [...]int{1,2,3,4,5}
	var arr5 = [...]int{1,2,3,4,5}
	var arr6 [5] int = [...]int{1,2,3,4,5}//位数不能超5

	fmt.Println(arr,arr1,arr2,arr3,arr4,arr5,arr6)

	for i2, i3 := range arr {
		fmt.Printf("key:%d,value:%d \n",i2,i3)
	}
	for _, v := range arr {
		fmt.Printf("value:%d \n",v)
	}

	arr7 := arr

	fmt.Println(arr7)

	arr8 := [13]int{1,45,1,5,5,5466,4,5,86,5,4,1,56}
	fmt.Println(arr8)

	fmt.Println(sort(arr8))



}

func sort(arr [13]int) ([13]int) {
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp;
			}
		}
	}
	return arr
}