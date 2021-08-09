package main

import "fmt"

type People struct {
	sex string
	height float32
	size float32
	age int
}


func main () {
	var age = 100
	var ppp = "aaa"
	var test int = 1000

	var res = fmt.Sprintf("%s,,,,%d",ppp,age)
	fmt.Println(res)
	fmt.Println(test+1000)

	fmt.Println("aa")
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	fmt.Println(balance[0])

	var bbb int = 10;
	fmt.Println(bbb)
	fmt.Println(&bbb)
	fmt.Println(bbb)

	var ccc *int
	fmt.Println(ccc)
	ccc = &bbb
	fmt.Println(ccc)
	fmt.Println(bbb)
	fmt.Println(nil)


	var lzh = People{sex: "男",height: 183.01,size:64.01,age: 1000}
	fmt.Println(lzh)

	var xiaoming People;
	xiaoming.sex = "女"
	xiaoming.height = 1.00

	fmt.Println(xiaoming)
	testfunc();

	var 一个 = "沙雕"
	fmt.Println(一个)
}

func testfunc() {
	fmt.Println("test")
}