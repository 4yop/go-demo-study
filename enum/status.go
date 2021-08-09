package enum

import "fmt"

const SHOW int = 1;
const HIDE int = 2;


func getStatus (status int) {
	var statusMap = map[int]string{SHOW:"显示",HIDE:"隐藏"};
	fmt.Println(statusMap);
}