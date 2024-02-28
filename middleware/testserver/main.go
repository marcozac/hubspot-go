package main

import (
	"log"
	"net/http"
	"os"

	"github.com/marcozac/hubspot-go/middleware"
)

// This program is a simple HTTP server that listens on port 3000.
// It is used to test the middleware.NewSignatureV3 function.
// It responds with a 200 OK to the /ok endpoint and uses the middleware.NewSignatureV3
// to verify the signature of the request to the /test-signature-v3 endpoint.
// The client secret is read from the HUBSPOT_CLIENT_SECRET environment variable.
//
// The request to the /test-signature-v3 endpoint must come from HubSpot, for example,
// by creating a workflow with a webhook action that sends a POST request to the server.
//
// NGROK is a good tool to expose the server to the internet for testing.
func main() {
	clientSecret := []byte(os.Getenv("HUBSPOT_CLIENT_SECRET"))
	mux := http.NewServeMux()
	handlerOk := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.Handle("/ok", handlerOk)
	mux.Handle("/test-signature-v3", middleware.NewSignatureV3(clientSecret, handlerOk))
	log.Fatal(http.ListenAndServe(":3000", mux)) //nolint:gosec
}
