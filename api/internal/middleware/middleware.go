package middleware

import "net/http"

func ContentTypeMiddleware(newxt http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		newxt.ServeHTTP(w, r)
	})
}
