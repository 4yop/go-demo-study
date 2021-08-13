package main

import "fmt"

type T struct{
	name string
}
func (t *T) F()  {
	t.name = "123"
}
var abc = 1000
func main()  {
	fmt.Println(abc)
	a, _ := test()
	fmt.Println(a)

	fmt.Println(test1())
	t := T{}
	t.F()

	fmt.Println(t)

	add1 := func(n int) (res int){
		return n + 1
	}
	fmt.Println(add1(1))

	test2(1)

	arr := []int{12,3,4,5}
	test2(arr...)
	start,end,arr1 := test3(arr...)
	fmt.Println(start,end,arr1)

	fmt.Println(getSum(1))
	fmt.Println(getSum(100))

	fmt.Println(opera(1,1000,add))

	fmt.Println("------------------------")
	func1 := closure()

	fmt.Println(func1())
	fmt.Println(func1())
	fmt.Println(func1())

	var aaa  = add
	fmt.Println(aaa(1,1))
}



func test () (int,string){
	return 1,"lzh"
}

func test1() (age int,name string){
	age =1
	name = "sdsdf"
	return
}


func test2(num...int)  {
	fmt.Println(abc)
	fmt.Println(num)
}

func test3(arr...int) (int,int,[]int){
	start := arr[0]
	end := arr[len(arr)-1]
	return start,end,arr
}

func getSum (n int) (int){
	if (n == 1) {
		return 1
	}
	return getSum(n-1) + n
}

func opera (n1 int,n2 int,jisuan func(int,int)(int)) (int){
	res := jisuan(n1,n2)
	return res;
}

func add(n1 int,n2 int)  (int){
	return int(n1 + n2)
}

func cut(n1 int,n2 int) (int){
	return int(n1 - n2)
}


func closure () (func()(int)){
	i := 0
	res := func() (int) {
		i++
		return i
	}
	return res
}