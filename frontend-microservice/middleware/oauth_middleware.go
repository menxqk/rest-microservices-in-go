package middleware

import (
	"net/http"

	"github.com/menxqk/rest-microservices-in-go/common/oauth"
)

func RequireOAuth(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		err := oauth.AuthenticateRequest(r)
		if err != nil {
			w.WriteHeader(err.Status())
			w.Write([]byte(err.Message()))
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}
