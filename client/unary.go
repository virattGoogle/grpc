// package main

// import (
// 	"context"
// 	pb "grpc/proto"
// 	"log"
// 	"time"
// )

// func callSayHello(client pb.GreetServiceClient) {

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

// 	defer cancel()

// 	res, err := client.SayHello(ctx, &pb.Noparam{})

// 	if err != nil {
// 		log.Fatalf("Unable to make functioncall %v", err)
// 	}

// 	 log.Printf("%v", res.Message)
// }

