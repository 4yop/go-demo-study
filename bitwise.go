package main

import "fmt"

func main () {

	a := 6
	b := 7
	//a = 6 ,110
	//b = 7 ,111
	fmt.Printf("a:%d,%b \n",a,a)
	fmt.Printf("b:%d,%b \n",b,b)
	//a&b
	/**
	a :110
	a :111
	& :110 a b a且b为1的为1
	| :111 a b a或b为1 的为 1
	^ :001 a b 相同的为0
	^&:001 a b 都为 1的清0
	 */

	fmt.Printf("a b:%d,%b \n",a&b,a&b)
	fmt.Printf("a|b:%d,%b \n",a|b,a|b)
	fmt.Printf("a^b:%d,%b \n",a^b,a^b)
	fmt.Printf("a^&b:%d,%b \n",a^b,a^b)
	fmt.Printf("a<<b:%d,%b \n",a<<b,a<<b)
	fmt.Printf("a>>b:%d,%b \n",a>>b,a>>b)


	c := 15646
	d := 65456

	fmt.Printf("c:%d,%b \n",c,c)
	fmt.Printf("d:%d,%b \n",d,d)



	r1 := c&d

	/**
	c :0011110100011110
	d :1111111110110000
	r1:0011110100010000   15632
	 */

	fmt.Println(r1)
}