package main

import (
	"fmt"
	"http2-server/pkg/core"
	"net/http"
)

func main()  {
	client, err := core.CreateClient()
	if err != nil {
	   panic(err)
	}
	request, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
	   panic(err)
	}
	//request.Header.Set("Authority", "hello-grpc.example.com")
	response, err := client.Do(request)
	if err != nil {
	   panic(err)
	}
	fmt.Printf("%#v", response)
}
