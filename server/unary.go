package main

import( "context"
pb "grpc/proto")

func(s *helloServer) SayHello(ctx context.Context,req *pb.Noparam)(*pb.HelloResponse,error){
	var a = "hello"
	return &pb.HelloResponse{
		Message:&a},nil
}