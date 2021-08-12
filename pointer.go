package main

import "fmt"

func main()  {
	a := 100

	fmt.Println(a)
	fmt.Printf("a=%d,的内存地址为%x \n",a,&a)

	var p *int
	p = &a
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)

	var p1 **int
	p1 = &p
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println(**p1)

	var arr = [5]int{1,2,3,4,5}
	var arr1 *[5]int
	arr1 = &arr
	fmt.Println(arr)
	fmt.Println(*&arr1)
	fmt.Println(arr1[0])



	var arr2 [5]* int
	for i := 0; i < len(arr2); i++ {
		arr2[i] = &arr[i]
	}
	fmt.Println(arr2)
	fmt.Println(arr2[0])
	fmt.Println(*arr2[0])
}