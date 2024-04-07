package sportstracker

import "net/http"

type sessionTokenRoundTripper struct {
	token string
	next  http.RoundTripper
}

func (s *sessionTokenRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Sttauthorization", s.token)

	query := r.URL.Query()
	query.Add("token", s.token)
	r.URL.RawQuery = query.Encode()

	// TODO: change this back to a "do" method as this doesn't work for some reason?
	return s.next.RoundTrip(r)
}
