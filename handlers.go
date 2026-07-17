package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
		return
	}
	writeJSON(w, http.StatusOK, ordersDTO)
	return

}

func GetIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // вытаскиваем id из пути/конкретный заказ клиента
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {

		writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "Неверный запрос")
		return
	}


	for _, current := range details {
		if current.ID == idInt {
			writeJSON(w, http.StatusOK, current)
			return
		}
	}
	writeError(w, http.StatusNotFound, "ORDER_NOT_FOUND", "Заказ не найден")
	return
}

func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
		return
	}

	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Неправильный запрос")
		return
	}

	if req.Login == "" || req.Password == "" {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "логин и пароль обязательны")
		return
	}

	if req.Login != "client@example.com" || req.Password != "password" {
		writeError(w, http.StatusUnauthorized, "INVALID_CREDENTIALS", "Неправильный логин или пароль")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"token": "fake-token"})
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
		return
	}
	idInt2, err := strconv.Atoi(id)

	if err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "Некорректный запрос")
		return
	}

	var found bool
	for _, new_current := range ordersDTO {
		if new_current.ID == idInt2 {
			found = true
			break
		}
	}
	if found == false {
		writeError(w, http.StatusNotFound, "ORDER_NOT_FOUND", "Заказ не найден")
		return
	}

	var comment Comment
	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)

	
	err = json.Unmarshal(bodyBytes, &comment) 
	if err != nil {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Некорректный JSON")
		return
	}

	
	if strings.TrimSpace(comment.Text) == "" {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Комментарий не должен быть пустым")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": "Комментарий отправлен. Данные обновятся после синхронизации."})
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
		return
	}
	writeJSON(w, http.StatusOK, profile)
}

func GetDocumentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
		return
	}
	id := r.PathValue("id")
	id2, err := strconv.Atoi(id)
	if err != nil {

		writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "Неверный запрос")
		return
	}


	var res bool
	for _, Order := range ordersDTO {

		if Order.ID == id2 {
			res = true
			break
		}
	}
	if res == false {
		writeError(w, http.StatusNotFound, "ORDER_NOT_FOUND", "Заказ не найден")
		return
	}

	
	filtered := make([]DocumentDTO, 0)
	for _, doc := range documents {
		if doc.OrderID == id2 && doc.VisibleForClient == true {
			filtered = append(filtered, doc)
		}
	}

	for _, doc := range documents {
		if doc.OrderID == id2 && !doc.VisibleForClient {
			writeError(w, http.StatusForbidden, "FORBIDDEN", "Документы скрыты")
			return
		}
	}

	writeJSON(w, http.StatusOK, filtered)
	return

}

func PostLogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"text": "logged out"})

}

func GetAuthHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
		return
	}
	// создал переменную в обработчике
	user := Auth{Name: "Ivan", Email: "i_ivanov@test.com"}
	writeJSON(w, http.StatusOK, user)
}
