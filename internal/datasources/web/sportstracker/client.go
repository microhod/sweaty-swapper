package sportstracker

import (
	"net/http"
)

type Client struct {
	client   *http.Client
	baseURL  string
	pageSize int
}

type ClientOptions struct {
	BaseURL      string
	SessionToken string
	PageSize     int
}

type ClientOptionFunc func(*ClientOptions)

var defaultClientOptions = ClientOptions{
	BaseURL:  "https://api.sports-tracker.com/apiserver",
	PageSize: 100,
}

func NewClient(client *http.Client, optionFuncs ...ClientOptionFunc) *Client {
	options := defaultClientOptions
	for _, o := range optionFuncs {
		o(&options)
	}
	addSessionTokenAuth(client, options.SessionToken)

	return &Client{
		client:  client,
		baseURL: options.BaseURL,
		pageSize: options.PageSize,
	}
}
