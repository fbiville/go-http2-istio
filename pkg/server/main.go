package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"http2-server/pkg/core"
	"log"
	"net/http"
)

func main() {
	server, err := core.CreateServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Handler = h2c.NewHandler(BasicHandler{}, &http2.Server{})
	//server.Handler = BasicHandler{}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type BasicHandler struct{}

func (BasicHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Proto)
	for k, v := range req.Header {
		fmt.Printf("%v - %v\n", k, v)
	}
	res.WriteHeader(200)
	res.Header().Add("Content-Type", "text/plain")
	_, err := res.Write([]byte("hello world"))
	if err != nil {
		log.Fatal(err)
	}
}
