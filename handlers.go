package main

import (
    "encoding/json"
    "io"
    "net/http"
    "strconv"
)

func GetOrderHandler(w http.ResponseWriter, r *http.Request){

		if r.Method != http.MethodGet{
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusMethodNotAllowed)	
			json.NewEncoder(w).Encode(ErrorResp{Code: "METHOD_NOT_ALLOWED", Message: "Метод не обнаружен"})
			return
		}

		
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// отправляет JSON
		json.NewEncoder(w).Encode(orders)
		
	}

func GetIdHandler(w http.ResponseWriter, r *http.Request){
		id := r.PathValue("id") 

	
		if r.Method != http.MethodGet{
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(ErrorResp{Code: "METHOD_NOT_ALLOWED", Message: "Метод не обнаружен"})
			return 
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResp{Code: "INVALID_REQUEST", Message: "Некорректный запрос"})
			return
		}
		
		

		for _, current := range orders {
			if current.ID == idInt {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				json.NewEncoder(w).Encode(current)
				return
			}
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(ErrorResp{Code: "ORDER_NOT_FOUND", Message: "Заказ не найден"})
			return 
	}

func PostLoginHandler(w http.ResponseWriter, r *http.Request){	
	if r.Method != http.MethodPost{
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResp{Code: "METHOD_NOT_ALLOWED", Message: "Запрашиваемый адрес не поддерживает HTTP метод"})
		return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(map[string]string{"token": "fake-token"})
	}
	


func PostCommentHandler(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(ErrorResp{Code: "METHOD_NOT_ALLOWED", Message: "Метод не поддерживается"})
			return
		}
	
		
		var comment Comment
		defer r.Body.Close()
		
		bodyBytes, err := io.ReadAll(r.Body)
		 
		
		err = json.Unmarshal(bodyBytes, &comment) 
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResp{Code: "VALIDATION_ERROR", Message: "Некорректный JSON"})
			return
		}

		
		if comment.Text == ""{
			json.NewEncoder(w).Encode(ErrorResp{Code: "VALIDATION_ERROR", Message: "Комментарий не должен быть пустым"})
			return
		}
		json.NewEncoder(w).Encode(SuccessResp{Code: "STATUS_200", Message: Операция прошла успешно!"})
	}

func GetProfileHandler(w http.ResponseWriter, r *http.Request){	
	if r.Method != http.MethodGet{
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResp{Code: "METHOD_NOT_ALLOWED", Message: "Метод не поддерживается"})
		return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(profile)
}
	

func GetDocumentsHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w). Encode(ErrorResp{Code: "METHOD_NOT_ALLOWED", Message: "Статус не доступен"})
		return 
	}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(documents)
}
func PostLogoutHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w). Encode(ErrorResp{Code: "METHOD_NOT_ALLOWED", Message: "Метод не поддерживается"})
	return
}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(SuccessResp{Message: "Logged out"})

}

func GetAuthHandler(w http.ResponseWriter, r *http.Request){
	

	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w). Encode(ErrorResp{Code: "METHOD_NOT_ALLOWED", Message: "Метод не поддерживается"})
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	user := Auth{Name: "Ivan", Email: "i_ivanov@test.com"}
	json.NewEncoder(w).Encode(user)
}
