package main

import (
	"fmt"
	pb "grpc/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

const port = ":8080"

type helloServer struct{
	pb.GreetServiceServer
}

func main(){

	lis,err := net.Listen("tcp",port) 

	if err != nil {
		log.Fatalf("unable to start server")
	}
    grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer,&helloServer{})

	log.Printf("Server started listening on %v",lis.Addr())

	if err:= grpcServer.Serve(lis); err !=nil {
		log.Fatalf("Failed to start %v" ,err)
	}

}

func (s *helloServer)SayHelloServerStream(req *pb.NamesList,stream pb.GreetService_SayHelloServerStreamServer) error {


	log.Printf("Got request with Names ,%v",req.Names)

	for _,name := range req.Names {

		a :=  fmt.Sprintf("Hello",name)
		res := pb.HelloResponse{
			Message:&a,
		}

		if err := stream.Send(&res);err!= nil {
			return err
		}
		time.Sleep(2 *time.Second )
	}
	return nil
}