package main

import (
	"fmt"
	"strings"
)

func main()  {
	s1 := "A"
	s2 := "a"
	//s3 := "A"
	s4 := ""
	s5 := "  "

	isEq1 := strings.EqualFold(s1,s2);
	fmt.Printf("%s == %s ,%t \n",s1,s2,isEq1)
	isEq2 := strings.EqualFold(s4,s5)
	fmt.Printf("%s == %s ,%t \n",s4,s5,isEq2)

	s7 := "qwe"
	s8 := "qw"
	isHasPrefix1 := strings.HasPrefix(s7,s8)
	fmt.Printf("%s 有前缀 %s ? %t \n",s7,s8,isHasPrefix1)

	isHasPrefix2 := strings.HasPrefix(s7,s4)
	fmt.Printf("%s 有前缀 %s ? %t \n",s7,s4,isHasPrefix2)

	isHasPrefix3 := strings.HasPrefix(s7,s5)
	fmt.Printf("%s 有前缀 %s ? %t \n",s7,s5,isHasPrefix3)

	s9 := "qwe"
	s10 := "we"
	isHasSuffix1 := strings.HasSuffix(s9,s10)
	fmt.Printf("%s 有后缀 %s ? %t \n",s9,s10,isHasSuffix1)

	isContains1 := strings.Contains(s9,s10)
	fmt.Printf("%s 有包含 %s ? %t \n",s9,s10,isContains1)

	isContains2 := strings.Contains(s9,s5)
	fmt.Printf("%s 有包含 %s ? %t \n",s9,s5,isContains2)

	isContains3 := strings.Contains(s9,s4)
	fmt.Printf("%s 有包含 %s ? %t \n",s9,s4,isContains3)

	s11 := "ABCDEFa"
	s12 := 'A'
	isContainsRune1 := strings.ContainsRune(s11,s12)
	fmt.Printf("%s 有包含UTF-8码值 %d ? %t \n",s11,s12,isContainsRune1)

	s13 := "123456789"
	s14 := "13579"
	isContainsAny1 := strings.ContainsAny(s13,s14)
	fmt.Printf("%s 有一样的字符和 %s ? %t \n",s13,s14,isContainsAny1)

	isContainsAny2 := strings.ContainsAny(s13,s4)
	fmt.Printf("%s 有一样的字符和 %s ? %t \n",s13,s4,isContainsAny2)

	isContainsAny3 := strings.ContainsAny(s13,s5)
	fmt.Printf("%s 有一样的字符和 %s ? %t \n",s13,s5,isContainsAny3)

	isContainsAny4 := strings.ContainsAny(s4,s5)
	fmt.Printf("%s 有一样的字符和 %s ? %t \n",s4,s5,isContainsAny4)
}