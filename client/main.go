package main

import (
	"context"
	pb "grpc/proto"
	"io"
	"log"
	"time"

	// "golang.org/x/text/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":8080"

func main(){
	conn,err := grpc.Dial("127.0.0.1"+port,grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err !=nil{
		log.Fatalf("Could not connect to the server")
	}
defer conn.Close()

 client := pb.NewGreetServiceClient(conn)

 names := pb.NamesList{
 	Names: []string{"Vinayak","Parul","Kian","Nono"},
 }

 //callSayHello(client)

 callSayHelloServerStreaming(client,names)



}

func callSayHelloServerStreaming(client pb.GreetServiceClient, names pb.NamesList) {
	
	log.Printf("Streaming has started")

	stream,err := client.SayHelloServerStream(context.Background(),&names)

	if err!=nil {
		log.Fatalf("could not send names %v",err)
	}

	for{
		message,err := stream.Recv()

		if err == io.EOF{
			break
		}

		if err != nil{
			log.Fatalf("Unable to stream %v",err)

		}
        
        log.Println(message)
	}
       log.Printf("Streaming Finished")
}




func callSayHello(client pb.GreetServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.SayHello(ctx, &pb.Noparam{})

	if err != nil {
		log.Fatalf("Unable to make functioncall %v", err)
	}

	 log.Printf("%v", res.Message)
}


func SayHelloServerStreaming(client *pb.GreetServiceClient,names *pb.NamesList){


}
