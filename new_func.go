package main

import "net/http"

func newRouter() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api/cabinet/auth/logout", AuthMiddleware(http.HandlerFunc(PostLogoutHandler)))
	mux.Handle("/api/cabinet/auth/me", AuthMiddleware(http.HandlerFunc(GetAuthHandler)))
	mux.Handle("/api/cabinet/orders", AuthMiddleware(http.HandlerFunc(GetOrderHandler)))
	mux.Handle("/api/cabinet/orders/{id}", AuthMiddleware(http.HandlerFunc(GetIdHandler)))
	mux.Handle("/api/cabinet/orders/{id}/comment", AuthMiddleware(http.HandlerFunc(PostCommentHandler)))
	mux.Handle("/api/cabinet/profile", AuthMiddleware(http.HandlerFunc(GetProfileHandler)))
	mux.Handle("/api/cabinet/orders/{id}/documents", AuthMiddleware(http.HandlerFunc(GetDocumentsHandler)))
	mux.HandleFunc("/api/cabinet/auth/login", PostLoginHandler)

	return mux
}
