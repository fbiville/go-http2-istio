package core

import (
	"crypto/tls"
	"golang.org/x/net/http2"
	"net/http"
)

func CreateClient() (*http.Client, error) {
	client := http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return &client, nil
}