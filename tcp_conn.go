package main

import (
	"fmt"
	"net"
)

func main()  {

	conn,err:= net.Dial("tcp","127.0.0.1:7684")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()


	_,err = conn.Write([]byte("are you ok?"))
	if err != nil {
		fmt.Println(err)
	}



}
