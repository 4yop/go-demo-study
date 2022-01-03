package main

import (
	"context"
	trippd "coolcar/proto/gen/go"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("127.0.0.1:8081",grpc.WithInsecure())
	if err != nil {
		log.Fatalln("grpc dial error: %v", err)
	}
	tsClient := trippd.NewTripServiceClient(conn)
	r, err := tsClient.GetTrip(context.Background(),
							&trippd.GetTripRequest{
								Id: "trip45678",
							})
	if err != nil {
		log.Fatalln("tsClient.GetTrip err: $v", err)
	}
	fmt.Println(r)
}
