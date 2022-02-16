package middleware

import "net/http"

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST,PATCH,DELETE,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("*"))
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}
