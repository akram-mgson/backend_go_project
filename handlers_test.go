// шаблон для тестов

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
	body := rr.Body.String()
	jsonBytes := []byte(body)
	var orders []OrderListItemDTO

	
	err = json.Unmarshal(jsonBytes, &orders)
	if err != nil {
		t.Fatalf("ошибка парсинга: %v", err)
	}
	if len(orders) != 3 {
		t.Fatalf("ожидается 3 заказа, получен %d", len(orders))

	}
	if orders[0].ID != 40364 {
		t.Fatalf("получен ID 40364 %d", orders[0].ID)
	}

	if orders[1].ID != 40365 {
		t.Fatalf("получен ID 40365 %d", orders[0].ID)
	}
	if orders[2].ID != 40366 {
		t.Fatalf("получен ID 40366 %d", orders[0].ID)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("expected 200, got %v", status)
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

	respBody := rr.Body.Bytes()
	var errResp ErrorResp
	err = json.Unmarshal(respBody, &errResp)
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "UNAUTHORIZED" {
		t.Errorf("expected UNAUTHORIZED, got %v", errResp.Code)
	}
	if errResp.Message != "Неверный токен" {
		t.Errorf("expected 'Неверный токен', got %v", errResp.Message)
	}

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %v", status)
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
	body := rr.Body.String()
	jsonBytes := []byte(body)
	var order OrderDetailsDTO

	err = json.Unmarshal(jsonBytes, &order)
	if err != nil {
		t.Fatalf("ошибка парсинга: %v", err)
	}
	if order.ID == 0 {
		t.Fatalf("ожидался заказ, но ответ пустой")
	}
	if order.ID != 40364 {
		t.Fatalf("получен ID 40364 %d", order.ID)
	}
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("expected 200, got %v", status)
	}

}

func TestGetOrderInvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/orders/abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	body := rr.Body.Bytes()
	var errResp ErrorResp
	err = json.Unmarshal(body, &errResp)
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "INVALID_REQUEST" {
		t.Errorf("expected INVALID_REQUEST, got %v", errResp.Code)
	}
	if errResp.Message != "Неверный запрос" {
		t.Errorf("expected 'Неверный запрос', got %v", errResp.Message)
	}
	if status := rr.Code; status != http.StatusBadRequest {
		t.Fatalf("expected 400, got %v", status)
	}

}

func TestGetOrderNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/orders/999", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	body := rr.Body.Bytes()
	var errResp ErrorResp
	err = json.Unmarshal(body, &errResp)
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "ORDER_NOT_FOUND" {
		t.Errorf("expected ORDER_NOT_FOUND, got %v", errResp.Code)
	}
	if errResp.Message != "Заказ не найден" {
		t.Errorf("expected 'Заказ не найден', got %v", errResp.Message)
	}
	if status := rr.Code; status != http.StatusNotFound {
		t.Fatalf("expected 404, got %v", status)
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

	respBody := rr.Body.Bytes()
	var errResp ErrorResp
	err = json.Unmarshal(respBody, &errResp)
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "UNAUTHORIZED" {
		t.Errorf("expected UNAUTHORIZED, got %v", errResp.Code)
	}
	if errResp.Message != "Неверный токен" {
		t.Errorf("expected 'Неверный токен', got %v", errResp.Message)
	}
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %v", status)
	}

}

func TestCreateCommentEmptyText(t *testing.T) {
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

	var comment MessageResponse

	err = json.Unmarshal(jsonBytes, &comment)
	if err != nil {
		t.Fatalf("ошибка парсинга: %v", err)
	}
	if comment.Message == "" {
		t.Fatalf("ожидается сообщение в ответе")
	}

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("expected 200, got %v", status)
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
		t.Fatalf("expected 400, got %v", status)
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

	respBody := rr.Body.Bytes()
	var errResp ErrorResp
	err = json.Unmarshal(respBody, &errResp)
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "UNAUTHORIZED" {
		t.Errorf("expected UNAUTHORIZED, got %v", errResp.Code)
	}
	if errResp.Message != "Неверный токен" {
		t.Errorf("expected 'Неверный токен', got %v", errResp.Message)
	}

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %v", status)
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

	body := rr.Body.String()
	jsonBytes := []byte(body)

	var profile Info
	err = json.Unmarshal(jsonBytes, &profile)
	if profile.Email == "" {
		t.Fatalf("ожидается email в ответе")
	}
	if profile.PhoneNumber == "" {
		t.Fatalf("ожиадется phone_number в ответе")
	}
	if profile.Manager == "" {
		t.Fatalf("ожидается manager в ответе")
	}
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("expected 200, got %v", status)
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

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %v", status)
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
		t.Fatalf("ошибка парсинга: %v", err)
	}
	if len(documents) == 0 {
		t.Fatalf("ожиадется список документов, но ответ пустой")
	}
	if documents[0].ID == 0 {
		t.Fatalf("документ должен иметь ID")
	}
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("expected 200, got %v", status)
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

	respBody := rr.Body.Bytes()
	var errResp ErrorResp
	err = json.Unmarshal(respBody, &errResp)
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "UNAUTHORIZED" {
		t.Errorf("expected UNAUTHORIZED, got %v", errResp.Code)
	}
	if errResp.Message != "Неверный токен" {
		t.Errorf("expected 'Неверный токен', got %v", errResp.Message)
	}
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %v", status)
	}

}

func TestLoginSuccess(t *testing.T) {
	body := strings.NewReader(`{"login":"client@example.com", "password":"password"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/auth/login", body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	jsonBytes := rr.Body.Bytes()

	var resp LoginResponse
	err = json.Unmarshal(jsonBytes, &resp)
	if resp.Token == "" {
		t.Fatalf("ожидается token в ответе")
	}

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("expected 200, got %v", status)
	}

}

func TestLoginInvalidPassword(t *testing.T) {
	body := strings.NewReader(`{"login":"client@example.com", "password":"wrong_password"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/auth/login", body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	
	router := newRouter()
	router.ServeHTTP(rr, req)

	respBody := rr.Body.Bytes()
	var errResp ErrorResp
	err = json.Unmarshal(respBody, &errResp)
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "INVALID_CREDENTIALS" {
		t.Errorf("expected INVALID_CREDENTIALS, got %v", errResp.Code)
	}
	if errResp.Message != "Неправильный логин или пароль" {
		t.Errorf("expected 'Неверный логин и пароль', got %v", errResp.Message)
	}

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %v", status)
	}
}

func TestLoginInvalidLogin(t *testing.T) {
	body := strings.NewReader(`{"login":"wrong@example.com", "password":"password"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/auth/login", body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	respBody := rr.Body.Bytes()
	var errResp ErrorResp
	err = json.Unmarshal(respBody, &errResp)
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "INVALID_CREDENTIALS" {
		t.Errorf("expected INVALID_CREDENTIALS, got %v", errResp.Code)
	}
	if errResp.Message != "Неправильный логин или пароль" {
		t.Errorf("expected 'Неверный логин и пароль', got %v", errResp.Message)
	}

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %v", status)
	}
}

func TestLoginEmptyCredentials(t *testing.T) {
	body := strings.NewReader(`{}`)
	req, err := http.NewRequest("POST", "/api/cabinet/auth/login", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	respBody := rr.Body.Bytes()
	var errResp ErrorResp
	err = json.Unmarshal(respBody, &errResp)
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "VALIDATION_ERROR" {
		t.Errorf("expected INVALID_CREDENTIALS, got %v", errResp.Code)
	}
	if errResp.Message != "логин и пароль обязательны" {
		t.Errorf("expected 'логин и пароль', got %v", errResp.Message)
	}

	if status := rr.Code; status != http.StatusBadRequest {
		t.Fatalf("expected 400, got %v", status)
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
	body := rr.Body.Bytes()

	var errResp ErrorResp
	err = json.Unmarshal(body, &errResp)
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}

	if errResp.Code != "ORDER_NOT_FOUND" {
		t.Fatalf("expected ORDER_NOT_FOUND, got %v", errResp.Code)
	}
	if errResp.Message != "Заказ не найден" {
		t.Fatalf("expected 'Заказ не найден', got %v", errResp.Message)
	}

	if status := rr.Code; status != http.StatusNotFound {
		t.Fatalf("expected 404, got %v", status)
	}

}


func TestAuthMeSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/auth/me", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("expected 200, got %v", status)
	}

	var auth Auth
	err = json.Unmarshal(rr.Body.Bytes(), &auth)
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}

	if auth.Name == "" {
		t.Errorf("ожидается name в ответе")
	}
	if auth.Email == "" {
		t.Errorf("ожидается email в ответе")
	}

}

func TestAuthMeUnauthorized(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/cabinet/auth/me", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %v", status)
	}

	var errResp ErrorResp
	err = json.Unmarshal(rr.Body.Bytes(), &errResp)
	if err != nil {
		t.Fatalf("Ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "UNAUTHORIZED" {
		t.Errorf("expected UNAUTHORIZED, got %v", errResp.Code)
	}
	if errResp.Message != "Неверный токен" {
		t.Errorf("expected 'Неверный токен', got %v", errResp.Message)
	}

}

func TestLogoutSuccess(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/cabinet/auth/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("expected 200, got %v", status)
	}

	var resp map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Ошибка парсинга JSON: %v", err)
	}
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if resp["text"] != "logged out" {
		t.Errorf("expected 'logged out', got %v", resp["text"])
	}

}

func TestLogoutUnauthorized(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/cabinet/auth/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %v", status)
	}
	var errResp ErrorResp
	err = json.Unmarshal(rr.Body.Bytes(), &errResp)
	if err != nil {
		t.Fatalf("Ошибка парсинга JSON: %v", err)
	}
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "UNAUTHORIZED" {
		t.Errorf("expected UNAUTHORIZED, got %v", errResp.Code)
	}
	if errResp.Message != "Неверный токен" {
		t.Errorf("expected 'Неверный токен', got %v", errResp.Message)
	}

}

func TestLogoutMethodNotAllowed(t *testing.T) {

	req, err := http.NewRequest("GET", "/api/cabinet/auth/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %v", status)
	}
	var errResp ErrorResp
	err = json.Unmarshal(rr.Body.Bytes(), &errResp)
	if err != nil {
		t.Fatalf("Ошибка парсинга JSON: %v", err)
	}
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "METHOD_NOT_ALLOWED" {
		t.Errorf("expected METHOD_NOT_ALLOWED, got %v", errResp.Code)
	}
	if errResp.Message != "Метод не поддерживается" {
		t.Errorf("expected 'Метод не поддерживается', got %v", errResp.Message)
	}

}

func TestAuthMeMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/cabinet/auth/me", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")
	rr := httptest.NewRecorder()
	router := newRouter()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %v", status)
	}
	var errResp ErrorResp
	err = json.Unmarshal(rr.Body.Bytes(), &errResp)
	if err != nil {
		t.Fatalf("Ошибка парсинга JSON: %v", err)
	}
	if err != nil {
		t.Fatalf("ошибка парсинга JSON: %v", err)
	}
	if errResp.Code != "METHOD_NOT_ALLOWED" {
		t.Errorf("expected METHOD_NOT_ALLOWED, got %v", errResp.Code)
	}
	if errResp.Message != "Метод не поддерживается" {
		t.Errorf("expected 'Метод не поддерживается', got %v", errResp.Message)
	}

}
