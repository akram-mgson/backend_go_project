package main

import (
    "encoding/json"
    "io"
    "net/http"
    "strconv"
)

func GetOrderHandler(w http.ResponseWriter, r *http.Request){

		if r.Method != http.MethodGet{
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Заказ не найден")
			return
		}
		writeJSON(w, http.StatusOK, orders)
		return

	}

func GetIdHandler(w http.ResponseWriter, r *http.Request){
		id := r.PathValue("id") 
		if r.Method != http.MethodGet{
			writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не обнаружен" )
			return 
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			
			writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "Некорректный запрос")
		
			return
		}
		
		

		for _, current := range orders {

			
			if current.ID == idInt {
				writeJSON(w, http.StatusOK, current)
				return
			}
		}
		writeError(w, http.StatusNotFound, "ORDER_NOT_FOUND", "Заказ не найден")
			return 
	}

func PostLoginHandler(w http.ResponseWriter, r *http.Request){	
	if r.Method != http.MethodPost{
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "HTTP-запрос не поддерживается")
		return
		}

		var login LoginRequest

		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil{ 
			writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Неправильный запрос")
			return
		}

		if login.Login != "client@example.com" || login.Password != "password"{
		writeError(w, http.StatusUnauthorized, "INVALID_CREDENTIALS", "Неправильный логин или пароль")
		return
		}

		writeJSON(w, http.StatusOK, map[string]string{"token": "fake-token"})
	}


func PostCommentHandler(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
		if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
		return
	}
	idInt2, err := strconv.Atoi(id)
	if err != nil{
			writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "Некорректный запрос")
			return
	}
	var found bool
	for _, new_current := range orders{
		if new_current.ID == idInt2{
			found = true
			break
		}
	}
	if found == false{
		writeError(w, http.StatusNotFound, "STATUS_NOT_FOUND", "Статус не найден")
		return
	}

	
	if r.Method != http.MethodPost{
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не найден")
		return
	}
	id = r.PathValue("id")

	idInt2, err = strconv.Atoi(id)

	if err != nil{
			writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "Неверный запрос")
		return
	}


	for _, new_current := range orders{
		if new_current.ID == idInt2{
			found = true
			break
		}
	}
	if found == false{
		writeError(w, http.StatusNotFound, "ORDER_NOT_FOUND", "Заказ не найден")
		return 
	}


		
		var comment Comment
		defer r.Body.Close()
		bodyBytes, err := io.ReadAll(r.Body)
		 
		
		err = json.Unmarshal(bodyBytes, &comment) 
		if err != nil{
			writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Некорректный JSON")
			return
		}

		
		if comment.Text == ""{
			writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "Комментарий не должен быть пустым")
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"message": "comment sent"})
}


func GetProfileHandler(w http.ResponseWriter, r *http.Request){	
	if r.Method != http.MethodGet{
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
		return
		}
		writeJSON(w, http.StatusOK, profile)
	}

func GetDocumentsHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Статус не доступен")
		return
	}
	id := r.PathValue("id")
	id2, err := strconv.Atoi(id)
	if err != nil{
	
		writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "Неверный запрос")
		return
	}
		
		var res bool
		for _, Order := range orders{
			
			
				if Order.ID == id2{
					res = true
					break
				} 
			}
			if res == false {
				writeError(w, http.StatusNotFound, "ORDER_NOT_FOUND", "Заказ не найден")
				return 
			}


				
			filtered := make([]Document, 0)
			for _, doc := range documents{
				if doc.OrderID == id2{
					filtered = append(filtered, doc)
				}
			}
			writeJSON(w, http.StatusOK, filtered)
			return
			
}
func PostLogoutHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
	writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
	return
}
	writeJSON(w, http.StatusOK, map[string]string{"text": "loggout"})

}

func GetAuthHandler(w http.ResponseWriter, r *http.Request){
	

	if r.Method != http.MethodGet{
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
		return
	}
	
	user := Auth{Name: "Ivan", Email: "i_ivanov@test.com"}
	writeJSON(w, http.StatusOK, user)
}
