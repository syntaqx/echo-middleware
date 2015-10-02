package remoteaddr

import "net/http"

// RemoteAddr http handler
type RemoteAddr struct {
}

const (
	TrueClientIP   = "True-Client-IP"   // Edge providers (such as Akamai)
	TrueRealIP     = "True-Real-IP"     // Proxies (like Nginx)
	XForwardedFor  = "X-Forwarded-For"  // External proxies (like an ELB or router)
	XOriginatingIP = "X-Originating-IP" // Defacto email header
)

// New creates a new RemoteAddr handler
func New() *RemoteAddr {
	r := &RemoteAddr{}
	return r
}

// Handler provides a http.HandlerFunc interface
func (ra *RemoteAddr) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ra.handleActualRequest(w, r)
		h.ServeHTTP(w, r)
	})
}

// HandlerFunc provides a Martini compatible handler interface
func (ra *RemoteAddr) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	ra.handleActualRequest(w, r)
}

// ServeHTTP provides a Negroni compatible interface
func (ra *RemoteAddr) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ra.handleActualRequest(w, r)
	next(w, r)
}

// handleActualRequest applies a filterr function that attempts to source the
// address of a remote client through potential network manipulation. Once
// attached, it replaces the `RemoteAddr` so it can be treated as trusted.
func (ra *RemoteAddr) handleActualRequest(w http.ResponseWriter, r *http.Request) {
	var ipAddress string
	var ipSources = []string{
		r.Header.Get(TrueClientIP),
		r.Header.Get(TrueRealIP),
		r.Header.Get(XForwardedFor),
		r.Header.Get(XOriginatingIP),
	}

	// Iterate the ipSources to determine the valid address
	for _, ip := range ipSources {
		if ip != "" {
			ipAddress = ip
			break
		}
	}

	// Set the RemoteAddr to the determine IP address, if detected
	if ipAddress != "" {
		r.RemoteAddr = ipAddress
	}
}
