package main

import (
	"context"
	trippd "coolcar/proto/gen/go"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn,err := grpc.Dial("127.0.0.1:8081")
	if err != nil {
		log.Fatalln("grpc dial error: %v",err)
	}
	tsClient := trippd.NewTripServiceClient(conn)
	tsClient.GetTrip(context.Background(),)
}
