package network

import "net/http"

func MethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			SetCorsHeaders(w)
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

func SetCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-user-id, X-organization-id")
	w.Header().Set("Access-Control-Allow-Origin", "https://dev4.env.teamtells.ru")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
