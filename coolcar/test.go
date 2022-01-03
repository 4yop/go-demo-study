package main

import (
	trippd "coolcar/proto/gen/go"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	a := trippd.Trip{
		Start:       "11",
		End:         "22",
		DurationSec: 33,
		FeeCent:     44,
	}
	var err error
	b, err := proto.Marshal(&a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%X\n", b)
	fmt.Println(string(b))

	var tarip trippd.Trip
	err = proto.Unmarshal(b, &tarip)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(&tarip)
}
