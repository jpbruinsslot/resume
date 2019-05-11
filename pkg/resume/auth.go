package resume

import "net/http"

// Authenticated is an extremely simple authentication
func Authenticated(r *http.Request, u, p string) bool {
	user, pass, ok := r.BasicAuth()
	if !ok {
		return false
	}

	if user == u && pass == p {
		return true
	} else {
		return false
	}
}
