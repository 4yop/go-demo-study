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

func main()  {
	rect := rect{1,3}

	show("rect",rect)

	c := circle{100}
	show("circle",c);

	show("test","test");
}

