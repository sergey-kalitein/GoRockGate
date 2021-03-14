package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
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

func HandlerPushNotifications(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	serviceId := vars["service"]
	log.Printf("Target Push Service ID: %s", serviceId)

	// Checking the content type
	contentType := request.Header.Get("Content-Type")
	if contentType != "application/json" {
		// TODO: cover with test
		SendOutError(responseWriter, "Unsupported content type", http.StatusUnsupportedMediaType)
		return
	}

	// Parsing the incoming Push Notification
	pushPacket := &RocketPushPacket{}
	pushBodyText, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(pushBodyText, pushPacket)
	if err != nil {
		// TODO: cover with test
		SendOutError(responseWriter, "Unable to unmarshal incoming push: "+err.Error(), http.StatusBadRequest)
		return
	}

	SendOutJSON(responseWriter, pushPacket, http.StatusOK)
}
