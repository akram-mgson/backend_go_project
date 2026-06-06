package main

import (
    "encoding/json"
    "net/http"
    
)


func writeJSON(w http.ResponseWriter, status int, data any) {
w.Header().Set("Content-Type", "application/json; charset=utf-8")
w.WriteHeader(status)
json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, code string, message string) {
writeJSON(w, status, ErrorResp{Code: code, Message: message})}
