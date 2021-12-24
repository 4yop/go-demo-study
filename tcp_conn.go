package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {

	conn,err:= net.Dial("tcp","127.0.0.1:7684")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	
	go func() {
		b := make([]byte,1024)
		for  {
			
		}
	}()


	b := make([]byte,1024)
	for{

		n,err := os.Stdin.Read(b)//读键盘内容
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b[:n]))

		conn.Write(b[:n])

	}


}
