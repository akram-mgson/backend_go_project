package main

import (
	"log"
	"net/http"

	
)

func main() {
	
	db, err := InitDB()
	if err != nil{
		log.Fatal("Нет подключения к БД", err)
	}
	defer db.Close()


	
	http.HandleFunc("/api/cabinet/auth/login", PostLoginHandler)
	http.HandleFunc("/api/cabinet/auth/logout", PostLogoutHandler)
	http.HandleFunc("/api/cabinet/auth/me", GetAuthHandler)

	// обертка handle
	http.Handle("/api/cabinet/orders", AuthMiddleware(http.HandlerFunc(GetOrderHandler)))
	http.Handle("/api/cabinet/orders/{id}", AuthMiddleware(http.HandlerFunc(GetIdHandler)))
	http.Handle("/api/cabinet/orders/{id}/comment", AuthMiddleware(http.HandlerFunc(PostCommentHandler)))
	http.Handle("/api/cabinet/profile", AuthMiddleware(http.HandlerFunc(GetProfileHandler)))
	http.Handle("/api/cabinet/orders/{id}/documents", AuthMiddleware(http.HandlerFunc(GetDocumentsHandler)))
	

	// работа и запуск сервера
	log.Println("запуск сервера http://localhost:8080")

	// запуск по порталу
	log.Fatal(http.ListenAndServe(":8080", nil))
}

	
