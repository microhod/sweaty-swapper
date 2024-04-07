package sportstracker

import (
	"net/http"
)

const apiserverURL = "https://api.sports-tracker.com/apiserver"

type Client struct {
	sessionToken string
	client       httpClient
}

type httpClient interface {
	Do(r *http.Request) (*http.Response, error)
}

func NewClient(sessionToken string) *Client {
	return &Client{
		sessionToken: sessionToken,
		client:       http.DefaultClient,
	}
}

func (c *Client) Do(r *http.Request) (*http.Response, error) {
	r.Header.Add("Sttauthorization", c.sessionToken)

	query := r.URL.Query()
	query.Add("token", c.sessionToken)
	r.URL.RawQuery = query.Encode()

	return c.client.Do(r)
}

func (c *Client) Workouts() *WorkoutClient {
	return &WorkoutClient{
		client:   c,
		pageSize: defaultWorkoutsPageSize,
	}
}
