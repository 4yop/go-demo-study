package main

import (
	"fmt"
	"net"
)

func main() {
	conn,err := net.Dial("tcp","127.0.0.1:7812")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conn)
	b := make([]byte,1024)
	for  {
		n,err := conn.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b[:n]))
	}

}
