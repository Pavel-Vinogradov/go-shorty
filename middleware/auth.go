package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}
		bearer := strings.TrimPrefix(auth, "Bearer")
		if bearer == "" {
			next.ServeHTTP(w, r)
			return
		}
		fmt.Println("bearer is", bearer)
	})
}
