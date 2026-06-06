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

			// сравнение заказа с тем, что пришло в запросе
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
		
		bodyLogin, err := io.ReadAll(r.Body)
		if err != nil{ 
		defer r.Body.Close()
		}
		err = json.Unmarshal(bodyLogin, &login)

		if login.Login != "client@example.com" || login.Password != "password"{
		writeError(w, http.StatusUnauthorized, "INVALID_CREDENTIALS", "Неправильный логин или пароль")
		return
		}
		// if err != nil{
		// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	json.NewEncoder(w).Encode(ErrorResp{Code: "VALIDATION_ERROR", Message: "Ошибка валидации"})
		// 	//json.NewEncoder(w).Encode(map[string]string{"token": "fake-token"})
		// 	return
		// }

		writeJSON(w, http.StatusOK, map[string]string{"token": "fake-token"})
	}


func PostCommentHandler(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Метод не поддерживается")
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

		// обращаюсь к полю Text структуры Comment
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
		// for _, doc := range documents{
		// 	if doc.OrderID == id2{
		// 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 		json.NewEncoder(w).Encode(documents)
		// 	}
			
		//}
		
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


				
			var filtered [] Document
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
