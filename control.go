
package main

import "fmt"

func main() {
	if 2>1 {
		fmt.Println("12243243r")
	}
	var a = 1;
	switch a {
	case 2:
		fmt.Println("sdfsdf")
	case 1:
		fmt.Println("sdfgsdgsdgsdgh")
		//页执行下一个case
		fallthrough
	case 3 :
		fmt.Println("qq")
		//中断
		break;
	default:
		fmt.Println("dfsgdfgdfffffffffffffff")
	}

}