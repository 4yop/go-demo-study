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
		for{

			n,err := os.Stdin.Read(b)//读键盘内容
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("键盘:",string(b[:n]))

			conn.Write(b[:n])

		}
	}()




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
