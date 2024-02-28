package util

import "net/http"

const (
	HeaderContentType   = "Content-Type"
	MIMEApplicationJSON = "application/json"
)

// SetJSONHeader sets the Content-Type header to application/json.
func SetJSONHeader(req *http.Request) {
	req.Header.Add(HeaderContentType, MIMEApplicationJSON)
}
