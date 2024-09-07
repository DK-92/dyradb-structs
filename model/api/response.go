package api

import "net/http"

const xLogIdHeader = "X-LOG-ID"

// Response
// API responses fit in this struct.
type Response struct {
	Headers     *http.Header    `json:"-"` // This field should not appear
	Information []string        `json:"information,omitempty"`
	Errors      []*errorMessage `json:"errors,omitempty"`
}

// IsNotValid
// Checks whether the response is valid. Workaround, because the current HTTP client doesn't handle errors very well.
func (r *Response) IsNotValid() bool {
	return len(r.Errors) > 0
}

func (r *Response) XLogIDHeader() string {
	if r.Headers == nil {
		return ""
	}

	return r.Headers.Get(xLogIdHeader)
}
