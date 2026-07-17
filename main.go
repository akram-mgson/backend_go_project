package main

import (
	"log"
	"net/http"
)

func main() {

	// работа и запуск сервера
	log.Println("запуск сервера http://localhost:8080")

	// запуск по порталу
	log.Fatal(http.ListenAndServe(":8080", newRouter()))
}
