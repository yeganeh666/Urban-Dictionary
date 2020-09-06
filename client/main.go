package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"simple-microservice/urban"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("^_^ Welcome To Urban Dictionary ^_^")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	defer cc.Close()
	client := urban.NewUrbanDCClient(cc)
	fmt.Println("What's Your Name? :)")
	name := ""
	fmt.Scanln(&name)
	req := &urban.NameRequest{Name: name}
	res, err := client.SendDefenitions(context.Background(), req)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Println("   * * * * * * * * * * * * * * * * * *")
	for {
		defenition, err := res.Recv()
		if err == io.EOF {

			break
		}
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}
		fmt.Println("Definition:\n", defenition.Definition)
		fmt.Println("Example:\n", defenition.Example)
		fmt.Println("WrittenOn:\n", defenition.WrittenOn)
		fmt.Println("   * * * * * * * * * * * * * * * * * *")
	}
	return
}
