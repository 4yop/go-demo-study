package main

import "fmt"

const A = 1
func main()  {
	const A = 2

	const A1,A2,A3 = 1,2,3;

	const (
		A4 int = 1
		A5 string = "qq"
	)

	const A6 = iota

	const (
		A7 = iota
		A8 = iota
	)

	fmt.Println(A)
	fmt.Println(A1)
	fmt.Println(A4,A5)
	fmt.Println(A6)
	fmt.Println(A7,A8)

	fmt.Println(A > 1)

	a := 2;
	b := 3;

	fmt.Println(a^b)

}