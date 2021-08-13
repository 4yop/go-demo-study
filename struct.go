
package main

import (
	"fmt";

)

type ListNode struct {
	val int
	next *ListNode

}

type People1 struct {
	Age int
	Name string
	Sex string
}


func main(){
	golang := People1{Age: 18,Name: "golang",Sex: "man"}

	xiaoming := People1{100,"xiaoming","woman"}





	fmt.Println(golang,xiaoming)

	arr := []int{1,2,3,4,5}
	fmt.Println(arr)

	p := new(People1)
	fmt.Println(*p)

	strc := struct {
		name string
		age int
	}{"lzh",1}
	fmt.Println(strc)
	var d dnf = dnf{"l",1}
	fmt.Println(d)

	ln := ListNode{}
	ln.val = 1



	fmt.Println(ln)
}

type dnf struct {
	string
	int

}

