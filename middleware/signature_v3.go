package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	HeaderSignatureV3      = "X-HubSpot-Signature-v3"
	HeaderRequestTimestamp = "X-HubSpot-Request-Timestamp"
)

// NewSignatureV3 returns a middleware that verifies the signature of the request
// using the HMAC-SHA256 algorithm with the client secret.
//
// The signature is expected to be in the X-HubSpot-Signature-v3 header.
func NewSignatureV3(clientSecret []byte, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if len(clientSecret) == 0 {
			http.Error(w, "missing client secret", http.StatusInternalServerError)
			return
		}
		sh := r.Header.Get(HeaderSignatureV3)
		if sh == "" {
			http.Error(w, "missing signature", http.StatusBadRequest)
			return
		}
		tsh := r.Header.Get(HeaderRequestTimestamp)
		if tsh == "" {
			http.Error(w, "missing timestamp", http.StatusBadRequest)
			return
		}

		// Reject the request if the timestamp is older than 5 minutes
		timestamp, err := strconv.ParseInt(tsh, 10, 64)
		if err != nil {
			http.Error(w, "invalid timestamp", http.StatusBadRequest)
			return
		}
		if time.Since(time.UnixMilli(timestamp)) > 5*time.Minute {
			http.Error(w, "timestamp too old", http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error reading body", http.StatusInternalServerError)
			return
		}

		uri := fmt.Sprintf("https://%s%s", r.Host, r.URL.String())
		signingString := r.Method + uri + string(body) + tsh
		h := hmac.New(sha256.New, clientSecret)
		h.Write([]byte(signingString))

		expectedSignature := base64.StdEncoding.EncodeToString(h.Sum(nil))
		if !hmac.Equal([]byte(sh), []byte(expectedSignature)) { // constant-time string comparison
			http.Error(w, "invalid signature", http.StatusForbidden)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)
	})
}
