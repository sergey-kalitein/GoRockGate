package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
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
	s, _ := json.MarshalIndent(payload, "", "  ")
	_, err := fmt.Fprint(responseWriter, string(s))
	if err != nil {
		log.Printf("SendOutJSON: %s\n", err.Error())
	} else {
		if IsLoggingPayloadEnabled == true {
			log.Println(color.New(color.FgHiBlue).Sprint(string(s)))
		}
	}
}

func SendOutError(w http.ResponseWriter, errorText string, errorCode int) {
	log.Print(color.New(color.BgRed, color.FgHiYellow).Printf("[ERROR] %s \n", errorText))
	SendOutJSON(w, ServiceError{ErrorText: errorText}, errorCode)
}
