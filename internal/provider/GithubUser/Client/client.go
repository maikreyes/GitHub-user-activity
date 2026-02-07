package client

import (
	"net/http"
	"time"
)

type Client struct {
	ApiUrl     string
	httpClient *http.Client
}

func NewClient(apiUrl string) *Client {
	return &Client{
		ApiUrl: apiUrl,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
