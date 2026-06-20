package main

import (
	"net/http"          
	"net/http/httptest" 
	"testing"           
	"strings"
)
func TestGetOrdersWithToken(t *testing.T){
	// t *testing.T 
	


	req, err := http.NewRequest("GET", "/api/cabinet/orders", nil)
	if err != nil{
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer fake-token")



	rr := httptest.NewRecorder()
	handler := AuthMiddleware(http.HandlerFunc(GetOrderHandler))
    handler.ServeHTTP(rr, req)


  
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("expected 200, got %v", status)
    }

	
}

func TestGetOrdersWithoutToken(t *testing.T){
	req, err := http.NewRequest("GET", "/api/cabinet/orders", nil)
	if err != nil{
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
  
	handler := AuthMiddleware(http.HandlerFunc(GetOrderHandler))
    handler.ServeHTTP(rr, req)


    if status := rr.Code; status != http.StatusUnauthorized {
        t.Errorf("expected 401, got %v", status)
    }

}



func TestGetIDWithToken(t *testing.T){
	req, err := http.NewRequest("GET", "/api/cabinet/orders/40364", nil)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	rr := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.Handle("/api/cabinet/orders/{id}", AuthMiddleware(http.HandlerFunc(GetIdHandler)))
	mux.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("expected 200, got %v", status)
    }

	
}


func TestGetIDWithNewToken(t *testing.T){
	req, err := http.NewRequest("GET", "/api/cabinet/orders/abc", nil)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	rr := httptest.NewRecorder()
	handler := AuthMiddleware(http.HandlerFunc(GetIdHandler))
    handler.ServeHTTP(rr, req)
	
    if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("expected 400, got %v", status)
    }

	
}

func TestGetIDWithSecondToken(t *testing.T){
	req, err := http.NewRequest("GET", "/api/cabinet/orders/999", nil)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	rr := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.Handle("/api/cabinet/orders/40364", AuthMiddleware(http.HandlerFunc(GetIdHandler)))
	mux.ServeHTTP(rr, req)


    if status := rr.Code; status != http.StatusNotFound{
        t.Errorf("expected 404, got %v", status)
    }

	
}





func TestGetIDWithoutToken(t *testing.T){
	req, err := http.NewRequest("GET", "/api/cabinet/orders/40364", nil)
	if err != nil{
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := AuthMiddleware(http.HandlerFunc(GetIdHandler))
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusUnauthorized {
        t.Errorf("expected 401, got %v", status)
    }

	
}




func TestPostCommentWithToken(t *testing.T){
	body := strings.NewReader(`{"text":"hello"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/orders/40364/comment", body)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.Handle("/api/cabinet/orders/{id}/comment", AuthMiddleware(http.HandlerFunc(PostCommentHandler)))
	mux.ServeHTTP(rr, req)

   
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("expected 200, got %v", status)
    }

	
}



func TestPostCommentWithNewToken(t *testing.T){
	body := strings.NewReader(`{"text":""}`)
	req, err := http.NewRequest("POST", "/api/cabinet/orders/40364/comment", body)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.Handle("/api/cabinet/orders/{id}/comment", AuthMiddleware(http.HandlerFunc(PostCommentHandler)))
	mux.ServeHTTP(rr, req)

   


   
    if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("expected 400, got %v", status)
    }

	
}






func TestPostCommentWithoutToken(t *testing.T){
	body := strings.NewReader(`{"text":"hello"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/orders/40364/comment", body)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := AuthMiddleware(http.HandlerFunc(PostCommentHandler))
    handler.ServeHTTP(rr, req)
   
    if status := rr.Code; status != http.StatusUnauthorized {
        t.Errorf("expected 401, got %v", status)
    }

	
}



func TestGetProfileWithToken(t *testing.T){
	req, err := http.NewRequest("GET", "/api/cabinet/profile/", nil)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")
	rr := httptest.NewRecorder()
	handler := AuthMiddleware(http.HandlerFunc(GetProfileHandler))
    handler.ServeHTTP(rr, req)


    if status := rr.Code; status != http.StatusOK {
        t.Errorf("expected 200, got %v", status)
    }

	
}


func TestGetProfileWithoutToken(t *testing.T){
	req, err := http.NewRequest("GET", "/api/cabinet/profile", nil)
	if err != nil{
		t.Fatal(err)
	}


	rr := httptest.NewRecorder()
	handler := AuthMiddleware(http.HandlerFunc(GetProfileHandler))
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusUnauthorized {
        t.Errorf("expected 401, got %v", status)
    }

	
}




func TestGetDocumentsWithToken(t *testing.T){
	req, err := http.NewRequest("GET", "/api/cabinet/orders/40366/documents", nil)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	rr := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.Handle("/api/cabinet/orders/{id}/documents", AuthMiddleware(http.HandlerFunc(GetDocumentsHandler)))
	mux.ServeHTTP(rr, req)





    if status := rr.Code; status != http.StatusOK {
        t.Errorf("expected 200, got %v", status)
    }

	
}


func TestGetDocumentsWithoutToken(t *testing.T){
	req, err := http.NewRequest("GET", "/api/cabinet/orders/40364/documents", nil)
	if err != nil{
		t.Fatal(err)
	}


	rr := httptest.NewRecorder()
	handler := AuthMiddleware(http.HandlerFunc(GetDocumentsHandler))
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusUnauthorized {
        t.Errorf("expected 401, got %v", status)
    }

	
}



func TestPostLoginWithToken(t *testing.T){
	body := strings.NewReader(`{"Login":"client@example.com", "Password":"password"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/auth/login", body)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer fake-token")

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostLoginHandler)
    handler.ServeHTTP(rr, req)


   
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("expected 200, got %v", status)
    }

	
}



func TestPostLoginWithoutToken(t *testing.T){
	body := strings.NewReader(`{"Login":"client@example.com", "Password":"password"}`)
	req, err := http.NewRequest("POST", "/api/cabinet/auth/login", body)
	if err != nil{
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostLoginHandler)
    handler.ServeHTTP(rr, req)


   
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("expected 401, got %v", status)
    }

	
}
