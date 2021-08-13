package main

import (
	"fmt"
	"math"
)

type geometry interface {
	tmp
}

type tmp interface {
	area() float32
	perim() float32
}

//矩形
type rect struct {
	len,wid float32
}

//计算矩形面积
func (r rect) area() float32{
	return r.len * r.wid
}

//计算矩形的周长
func (r rect) perim() float32{
	return 2 * (r.len + r.wid)
}

type circle struct {
	radius float32
}

func (c circle) area() float32{
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float32{
	return 2 * math.Pi * c.radius
}


func show (name string,param interface{}) {

	switch param.(type) {
	case geometry:
		// 类型断言
		fmt.Printf("area of %v is %v \n", name, param.(geometry).area())
		fmt.Printf("perim of %v is %v \n", name, param.(geometry).perim())
	default:
		fmt.Println("wrong type!")
	}
}
type s1 int
func main()  {
	//rect := rect{1,3}
	//
	//show("rect",rect)

	c := circle{100}
	show("circle",c);

	show("test","test");


	map1 := make(map[interface{}]interface{})

	map1[1] = 2
	map1["p"] = "q"

	fmt.Println(map1)

	var aaa interface{};

	aaa = 1

	aaa = "q"

	fmt.Println(aaa)
	var s s1 = 1
	fmt.Println(s)
}

type myfun func(int,int)(int)

func testq(a int, b int) (myfun) {
	return func(i2 int, i22 int) int {
		return i2 + i22
	}
}

//起别名

type sdsdrdf = int