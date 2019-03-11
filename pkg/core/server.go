package core

import (
	"golang.org/x/net/http2"
	"net/http"
)

func CreateServer() (*http.Server, error) {
	var server http.Server
	server.Addr = ":8080"
	err := http2.ConfigureServer(&server, nil)
	if err != nil {
		return nil, err
	}
	return &server, nil
}
