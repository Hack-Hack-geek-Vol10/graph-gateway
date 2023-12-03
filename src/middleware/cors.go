package middleware

import "net/http"

func SetCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", authTokenKey)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	})
}
