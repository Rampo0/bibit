package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"play.com/rpc_client/rpc/moviepb"
)

func doUnaryDetail(client moviepb.MovieServiceClient) {
	req := &moviepb.DetailRequest{
		Id: "tt4853102",
	}

	res, err := client.Detail(context.Background(), req)
	if err != nil {
		log.Printf("Error while calling detail request : %v", err)
		return
	}

	log.Printf("Response from MovieRPCService : %v", res)
}

func doUnarySearch(client moviepb.MovieServiceClient) {
	req := &moviepb.SearchRequest{
		Search: "Batman",
		Page:   "2",
	}

	res, err := client.Search(context.Background(), req)
	if err != nil {
		log.Printf("Error while calling detail request : %v", err)
		return
	}
	log.Printf("Response from MovieRPCService : %v", res)
}

func main() {
	opts := grpc.WithInsecure()

	clientDial, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}

	defer clientDial.Close()

	client := moviepb.NewMovieServiceClient(clientDial)
	doUnaryDetail(client)
	doUnarySearch(client)
}
