package main

import "fmt"

//内存地址指向的值,改内存地址对应的,值就变了(不管在哪)
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


	s1 := 1
	qqq(&s1)
	fmt.Println(s1)
	s2 := []int{1,3,5}
	bbb(&s2)
	fmt.Println(s2)

	ccc(arr2)
	fmt.Println(arr2)
	fmt.Println(arr)
}

func qqq(n *int)  {
	//fmt.Println(n)
	*n = 2
}

func bbb(arr *[]int) {
	p := *arr
	//fmt.Println(p)
	//fmt.Println(p[0])
	p[0] = 10086
}

func ccc(arr [5]*int) {
	fmt.Println(arr)
	p := *arr[0]
	fmt.Println(p)
	*arr[0] = 999
}