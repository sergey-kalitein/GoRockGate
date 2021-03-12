package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ContentTypeJson = "application/json"
)

type RocketPushPacket struct {
	Token   string `json:"token"`
	Options struct {
		CreatedAt string `json:"createdAt"` //"2021-03-12T10:09:17.925Z",
		CreatedBy string `json:"createdBy"` // <SERVER>
		Sent      bool   `json:"sent"`      // false
		Sending   int    `json:"sending"`   // 0
		From      string `json:"from"`      // "push"
		Title     string `json:"title"`     // "@sg"
		Text      string `json:"text"`      // "This is a push test message"
		UserId    string `json:"userId"`    // "gR6Hhq5aEDdGswSQY",
		Sound     string `json:"sound"`     // "default",
		Apn       struct {
			Text string `json:"text"` // "@sg:\nThis is a push test message"
		} `json:"apn"`
		SiteURL  string `json:"site_url"` // "https://sg.workspee.chat"
		Topic    string `json:"topic"`    // "com.app.collaborative.chat",
		UniqueId string `json:"uniqueId"` // "no33sYn6N2fb8JNXm"
	} `json:"options"`
}

type ServiceError struct {
	ErrorText string `json:"error_text"`
}

func handlerPushNotifications(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	serviceId := vars["service"]
	log.Printf("Target Push Service ID: %s", serviceId)

	// Checking the content type
	contentType := request.Header.Get("Content-Type")
	if contentType != "application/json" {
		// TODO: cover with test
		sendOutError(responseWriter, "Unsupported content type", http.StatusUnsupportedMediaType)
		return
	}

	// Parsing the incoming Push Notification
	pushPacket := &RocketPushPacket{}
	pushBodyText, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(pushBodyText, pushPacket)
	if err != nil {
		// TODO: cover with test
		sendOutError(responseWriter, "Unable to unmarshal incoming push: "+err.Error(), http.StatusBadRequest)
		return
	}

	sendOutJSON(responseWriter, pushPacket, http.StatusOK)
}

func sendOutJSON(responseWriter http.ResponseWriter, payload interface{}, errorCode int) {
	responseWriter.Header().Set("Content-Type", ContentTypeJson)
	responseWriter.WriteHeader(errorCode)
	s, _ := json.Marshal(payload)
	_, err := fmt.Fprintf(responseWriter, "%s", string(s))
	if err != nil {
		log.Printf("sendOutJSON: %s\n", err.Error())
	}
}

func sendOutError(w http.ResponseWriter, errorText string, errorCode int) {
	log.Printf("[ERROR] %s \n", errorText)
	sendOutJSON(w, ServiceError{ErrorText: errorText}, errorCode)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/push/{service:[a-z0-9]+}/send", handlerPushNotifications)
	http.Handle("/", router)

	fmt.Println("Push Gateway is listening...")
	http.ListenAndServe(":8181", nil)
}
