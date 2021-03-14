package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	ContentTypeJson = "application/json"
)

type ServiceError struct {
	ErrorText string `json:"error_text"`
}

func SendOutJSON(responseWriter http.ResponseWriter, payload interface{}, errorCode int) {
	responseWriter.Header().Set("Content-Type", ContentTypeJson)
	responseWriter.WriteHeader(errorCode)
	s, _ := json.Marshal(payload)
	_, err := fmt.Fprint(responseWriter, string(s))
	if err != nil {
		log.Printf("SendOutJSON: %s\n", err.Error())
	}
}

func SendOutError(w http.ResponseWriter, errorText string, errorCode int) {
	log.Printf("[ERROR] %s \n", errorText)
	SendOutJSON(w, ServiceError{ErrorText: errorText}, errorCode)
}
