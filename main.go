package main

import (
	
    "log"
    "net/http"
)

func main(){
	// функция, котор занимается обработкой эндпоинта


	// иными словами - обработка запросов
	http.HandleFunc("/api/cabinet/orders", GetOrderHandler)
	http.HandleFunc("/api/cabinet/orders/{id}", GetIdHandler)
	http.HandleFunc("/api/cabinet/auth/login", PostLoginHandler)
	http.HandleFunc("/api/cabinet/orders/{id}/comment", PostCommentHandler)
	http.HandleFunc("/api/cabinet/profile", GetProfileHandler)

	
	// работа и запуск сервера
	log.Println("запуск сервера http://localhost:8080")

	// запуск по порталу
	log.Fatal(http.ListenAndServe(":8080", nil))
}