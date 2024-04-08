package sportstracker

import "net/http"

type sessionTokenRoundTripper struct {
	token string
	next  http.RoundTripper
}

func addSessionTokenAuth(client *http.Client, sessionToken string) {
	transport := client.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	client.Transport = &sessionTokenRoundTripper{
		token: sessionToken,
		next:  transport,
	}
}

func (s *sessionTokenRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Sttauthorization", s.token)

	query := r.URL.Query()
	query.Add("token", s.token)
	r.URL.RawQuery = query.Encode()

	return s.next.RoundTrip(r)
}
