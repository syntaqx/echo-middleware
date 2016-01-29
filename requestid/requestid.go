// Package requestid implements an http handler that assigns a randomly
// generated id to each request.
//
// Example:
//
//     package main
//
//     import (
//         "fmt"
//         "net/http"
//
//         "github.com/syntaqx/echo-middleware/requestid"
//     )
//
//     func main() {
//         mux := http.NewServeMux()
//         rid := requestid.New()
//
//         mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//             fmt.Fprintf(w, "%s is %s", rid.HeaderKey, r.Header.Get(rid.HeaderKey))
//         })
//
//         http.ListenAndServe(":8080", rid.Handler(mux))
//     }
//
package requestid

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

// XRequestID is the default header the request id is set to.
const XRequestID = "X-Request-Id"

// GenerateFunc is the func used by the handler to generate the random id.
type GenerateFunc func() (string, error)

// RequestID is an http handler that generates an id and assigns it to a request
// header.
type RequestID struct {
	Generate  GenerateFunc
	HeaderKey string
}

// New returns a new RequestID handler instance.
func New() *RequestID {
	m := &RequestID{
		Generate:  generateID,
		HeaderKey: "undefined",
	}

	m.SetHeaderKey(XRequestID)

	return m
}

// SetHeaderKey sets the header key the request id will be assigned to.
func (m *RequestID) SetHeaderKey(key string) {
	m.HeaderKey = http.CanonicalHeaderKey(key)
}

// SetGenerate sets the generate function the request id will be generated from.
func (m *RequestID) SetGenerate(f GenerateFunc) {
	m.Generate = f
}

// Handler implements the http.Handler interface and can be registered to serve
// a particular path or subtree in a HTTP server.
func (m *RequestID) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.handleActualRequest(w, r)
		h.ServeHTTP(w, r)
	})
}

// HandlerFunc provides a Martini compatible handler.
func (m *RequestID) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	m.handleActualRequest(w, r)
}

// ServeHTTP provides a Negroni compatible handler.
func (m *RequestID) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	m.handleActualRequest(w, r)
	next(w, r)
}

func (m *RequestID) handleActualRequest(w http.ResponseWriter, r *http.Request) {
	id, err := m.Generate()
	if err == nil {
		r.Header.Set(m.HeaderKey, id)
	}
}

func generateID() (string, error) {
	r := make([]byte, 12)
	_, err := rand.Read(r)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(r), nil
}
