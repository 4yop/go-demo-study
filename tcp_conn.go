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
		for  {

		}
	}()

	_,err = conn.Write([]byte("are you ok?"))
	if err != nil {
		fmt.Println(err)
	}

	for{
		str := make([]byte,1024)
		os.Stdin.Read(str)
	}


}
