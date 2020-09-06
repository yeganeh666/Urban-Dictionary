package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"simple-microservice/model"
	"simple-microservice/urban"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) SendDefenitions(request *urban.NameRequest, response urban.UrbanDC_SendDefenitionsServer) error {
	definition, err := find(request.Name)
	if err != nil {
		return err
	}
	for i, _ := range definition.List {
		res := convertToRes(i, definition)
		err = response.Send(res)
		if err != nil {
			return err
		}
	}
	return nil
}
func find(name string) (*model.MyJsonName, error) {
	url := fmt.Sprintf("https://mashape-community-urban-dictionary.p.rapidapi.com/define?term=%v", name)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-host", "mashape-community-urban-dictionary.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "23134bd254msh7f073503080cdc6p198070jsne6699a066dc5")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var definetions *model.MyJsonName
	json.NewDecoder(res.Body).Decode(&definetions)
	return definetions, nil
}
func convertToRes(n int, name *model.MyJsonName) *urban.NameResponse {
	return &urban.NameResponse{
		Definition: name.List[n].Definition,
		Example:    name.List[n].Example,
		WrittenOn:  name.List[n].WrittenOn,
	}

}
func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("Server is listening on %v...", address)
	s := grpc.NewServer()
	urban.RegisterUrbanDCServer(s, &server{})
	s.Serve(lis)
}
