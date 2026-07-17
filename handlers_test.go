package main

import (
	"encoding/json"
	"net/http"          
	"net/http/httptest" 
	"strings"
	"testing" 
)

func TestGetOrdersWithToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/orders", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer fake-token")


	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)
	rr.Body.String()
	body := rr.Body.String()
	jsonBytes := []byte(body)
	var orders []OrderListItemDTO

	err = json.Unmarshal(jsonBytes, &orders)
	if err != nil {
		t.Errorf("ошибка парсинга: %v", err)
	}
	if len(orders) == 0 {
		t.Errorf("ожидается список заказов, но ответ пустой")
	}
	if orders[0].ID != 40364 {
		t.Errorf("получен ID 40364 %d", orders[0].ID)
	}

	if orders[1].ID != 40365 {
		t.Errorf("получен ID 40365 %d", orders[0].ID)
	}
	if orders[2].ID != 40366 {
		t.Errorf("получен ID 40366 %d", orders[0].ID)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected 200, got %v", status)
	}

}

func TestGetOrdersWithoutToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/orders", nil)
	if err != nil {
		t.Fatal(err)
	}

	
	rr := httptest.NewRecorder()


	router := newRouter()
	router.ServeHTTP(rr, req)

	
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("expected 401, got %v", status)
	}

}

func TestGetIDWithToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/orders/40364", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")


	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	rr.Body.String()
	body := rr.Body.String()
	jsonBytes := []byte(body)
	var order OrderDetailsDTO

	err = json.Unmarshal(jsonBytes, &order)
	if err != nil {
		t.Errorf("ошибка парсинга: %v", err)
	}
	if order.ID == 0 {
		t.Errorf("ожидался заказ, но ответ пустой")
	}

	if order.ID != 40364 {
		t.Errorf("получен ID 40364 %d", order.ID)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected 200, got %v", status)
	}

}

func TestGetIDWithNewToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/orders/abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected 400, got %v", status)
	}

}

func TestGetIDWithSecondToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/orders/999", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	rr := httptest.NewRecorder()

	router := newRouter()
	router.ServeHTTP(rr, req)

	
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("expected 404, got %v", status)
	}

}

func TestGetIDWithoutToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/orders/40364", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)


	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("expected 401, got %v", status)
	}

}

func TestPostCommentWithToken(t *testing.T) {
	body := strings.NewReader(`{"text":"hello"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/orders/40364/comment", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	jsonBytes := rr.Body.Bytes()

	var comment PostComm

	err = json.Unmarshal(jsonBytes, &comment)
	if err != nil {
		t.Errorf("ошибка парсинга: %v", err)
	}
	if comment.Message == "" {
		t.Errorf("ожидается сообщение в ответе")
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected 200, got %v", status)
	}

}

func TestPostCommentWithNewToken(t *testing.T) {
	body := strings.NewReader(`{"text":""}`)
	req, err := http.NewRequest("POST", "/api/cabinet/orders/40364/comment", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected 400, got %v", status)
	}

}

func TestPostCommentWithoutToken(t *testing.T) {
	body := strings.NewReader(`{"text":"hello"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/orders/40364/comment", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("expected 401, got %v", status)
	}

}

func TestGetProfileWithToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/profile", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	rr.Body.String()
	body := rr.Body.String()
	jsonBytes := []byte(body)

	var profile Info
	err = json.Unmarshal(jsonBytes, &profile)
	if profile.Email == "" {
		t.Errorf("ожидается email в ответе")
	}
	if profile.PhoneNumber == "" {
		t.Errorf("ожиадется phone_number в ответе")
	}
	if profile.Manager == "" {
		t.Errorf("ожидается manager в ответе")
	}
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected 200, got %v", status)
	}

}

func TestGetProfileWithoutToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/profile", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	// 3. Проверить статус что он 200
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("expected 401, got %v", status)
	}

}

func TestGetDocumentsWithToken(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/cabinet/orders/40366/documents", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	jsonBytes := rr.Body.Bytes()

	var documents []DocumentDTO
	err = json.Unmarshal(jsonBytes, &documents)
	if err != nil {
		t.Errorf("ошибка парсинга: %v", err)
	}
	if len(documents) == 0 {
		t.Errorf("ожиадется список документов, но ответ пустой")
	}
	if documents[0].ID == 0 {
		t.Errorf("документ должен иметь ID")
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected 200, got %v", status)
	}

}

func TestGetDocumentsWithoutToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/orders/40364/documents", nil)
	if err != nil {
		t.Fatal(err)
	}

	
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	// 3. Проверить статус что он 200
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("expected 401, got %v", status)
	}

}

func TestPostLoginWithToken(t *testing.T) {
	body := strings.NewReader(`{"Login":"client@example.com", "Password":"password"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/auth/login", body)
	if err != nil {
		t.Fatal(err)
	}
	// заголовок
	req.Header.Set("Authorization", "Bearer fake-token")

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	jsonBytes := rr.Body.Bytes()

	var resp LoginResponce
	err = json.Unmarshal(jsonBytes, &resp)
	if resp.Token == "" {
		t.Errorf("ожидается token в ответе")
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected 200, got %v", status)
	}

}

func TestPostLoginWithoutToken(t *testing.T) {
	body := strings.NewReader(`{"Login":"client@example.com", "Password":"password"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/auth/login", body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected 401, got %v", status)
	}

}

func TestPostLoginInvalidPassword(t *testing.T) {
	body := strings.NewReader(`{"Login":"client@example.com", "Password":"wrong_password"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/auth/login", body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("expected 401, got %v", status)
	}
}

func TestPostLoginInvalidLogin(t *testing.T) {
	body := strings.NewReader(`{"Login":"wrong@example.com", "Password":"password"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/auth/login", body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("expected 401, got %v", status)
	}
}

func TestPostLoginEmptyBody(t *testing.T) {
	body := strings.NewReader(`{}`)
	req, err := http.NewRequest("POST", "/api/cabinet/auth/login", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected 400, got %v", status)
	}

}

func TestGetDocumentsOrderNotFound(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/cabinet/orders/99999/documents", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("expected 404, got %v", status)
	}

}

func TestGetDocumentsNoAccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/orders/40365/documents", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("expected 403, got %v", status)
	}

}
