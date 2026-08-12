package main

import "net/http"

// объявление функции, след обработчик, котор вызывается после проверки
func AuthMiddleware(next http.Handler) http.Handler {
	// handlerfunc - функция, котор превращает обычную функцию в http.handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// вызов обработчкика, передает ему и w, и r

		// Если он не равен "Bearer fake-token" → отправить ошибку 401 и выйти.
		// Если равен → вызвать next.ServeHTTP(w, r).

		// Получить заголовок Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer fake-token" {
			writeError(w, http.StatusUnauthorized, "UNAUTHORIZED", "Неверный токен")
			return
		}
		// Если равен → вызвать next.ServeHTTP(w, r).

		next.ServeHTTP(w, r)

	})
}
