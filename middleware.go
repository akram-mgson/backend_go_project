package main

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer fake-token"{
			writeError(w, http.StatusUnauthorized, "UNAUTHORIZED", "Неверный токен")
			return
		}
			next.ServeHTTP(w,r)
	
			


	})
}

