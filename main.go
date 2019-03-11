package main

import (
	"http2-server/pkg/core"
	"log"
	"net/http"
)

func main() {
	server, err := core.CreateServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Handler = BasicHandler{}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type BasicHandler struct{}

func (BasicHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	res.Header().Add("Content-Type", "text/plain")
	_, err := res.Write([]byte("hello world"))
	if err != nil {
		log.Fatal(err)
	}
}
