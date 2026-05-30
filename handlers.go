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
				// отправка ошибки 405
			json.NewEncoder(w).Encode(ErrorResp{Code: "METHOD_NOT_ALLOWED", Message: "Метод не обнаружен"})
			return
		}

		// заголовок ответа
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// отправляет JSON
		json.NewEncoder(w).Encode(orders)
		
	}

func GetIdHandler(w http.ResponseWriter, r *http.Request){
		id := r.PathValue("id") // вытаскиваем id из пути/конкретный заказ клиента

		// проверка метода, совпадает ли запрос с ожидаемым?
		if r.Method != http.MethodGet{
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusMethodNotAllowed)
			// отправка ошибки 405
			json.NewEncoder(w).Encode(ErrorResp{Code: "METHOD_NOT_ALLOWED", Message: "Метод не обнаружен"})
		
			return 
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResp{Code: "INVALID_REQUEST", Message: "Некорректный запрос"})
			// если мы вводим "abc" вместо уникального номера
		
			return
		}
	
		
	
		//fmt.Printf("ID стал числом: %d\n", idInt) - исключаем тк вывод сообщение
		
		

		for _, current := range orders {

			// сравнение заказа с тем, что пришло в запросе
			if current.ID == idInt {

				// если совпадает, то отправляем json - выход
				// в противном случае - продолжаем цикл

				w.Header().Set("Content-Type", "application/json; charset=utf-8")

				json.NewEncoder(w).Encode(current)
				return
			}
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)

		// замена http или переход с http на json
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
			// установка заголовки
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResp{Code: "VALIDATION_ERROR", Message: "ошибка валидации"})
		return
	}
	
		
		var comment Comment
		//закрытие сетевого соединения после выполнения http-запроса
		// r.Body - тело запроса клиента,что он отправил
		defer r.Body.Close()
		// чтение всех данных из потока до тех пор пока не произойдет ошибка
		// без этой функции я не получу данные от клиента, закрывается чтоб не было утечки
		bodyBytes, err := io.ReadAll(r.Body)
		 
		// парсинг JSON, без него не прочитать что прислал запрос
		err = json.Unmarshal(bodyBytes, &comment) // передаю адрес переменной
		if err != nil{
			json.NewEncoder(w).Encode(ErrorResp{Code: "VALIDATION_ERROR", Message: "ошибка валидации!"})
			return
		}

		// обращаюсь к полю Text структуры Comment
		if comment.Text == ""{
			json.NewEncoder(w).Encode(ErrorResp{Code: "VALIDATION_ERROR", Message: "ошибка валидации!"})
			return
		}
		json.NewEncoder(w).Encode(SuccessResp{Code: "STATUS_200", Message: "операция прошла успешно!"})
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
	// создал переменную в обработчике
	user := Auth{Name: "Ivan", Email: "i_ivanov@test.com"}
	json.NewEncoder(w).Encode(user)
}
