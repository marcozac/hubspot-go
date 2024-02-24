package util

import "net/http"

const (
	HeaderContentType   = "Content-Type"
	MIMEApplicationJSON = "application/json"
)

func SetJSONHeader(req *http.Request) {
	req.Header.Add(HeaderContentType, MIMEApplicationJSON)
}
