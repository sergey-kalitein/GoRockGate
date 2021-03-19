package controllers

import (
	"encoding/json"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/tbalthazar/onesignal-go"
	"io/ioutil"
	"log"
	"net/http"
	"rockgate/app"
	"rockgate/models"
)

func SendPushNotification(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	serviceId := vars["service"]
	log.Printf("Target Push Service ID: %s", serviceId)

	// Checking the content type
	contentType := request.Header.Get("Content-Type")
	if contentType != "application/json" {
		// TODO: cover with test
		app.SendOutError(responseWriter, "Unsupported content type", http.StatusUnsupportedMediaType)
		return
	}

	// Parsing the incoming Push Notification
	pushPacket := &models.RocketPushPacket{}
	pushBodyText, _ := ioutil.ReadAll(request.Body)
	err := json.Unmarshal(pushBodyText, pushPacket)
	if err != nil {
		// TODO: cover with test
		app.SendOutError(responseWriter, "Unable to unmarshal incoming push: "+err.Error(), http.StatusBadRequest)
		return
	} else {
		if app.IsLoggingPayloadEnabled() == true {
			log.Print(color.New(color.FgHiBlue).Printf("Incoming push message body:\n %s\n", pushBodyText))
		}
	}

	notificationResponse, err := processOneSignalNotification(*pushPacket)
	if err != nil {
		app.SendOutError(responseWriter, "Unable to create notification: "+err.Error(), http.StatusBadRequest)
	} else {
		app.SendOutJSON(responseWriter, notificationResponse, http.StatusOK)
	}
}

func processOneSignalNotification(pushPacket models.RocketPushPacket) (*onesignal.NotificationCreateResponse, error) {
	// Find an App
	foundApp, err := app.Services.OneSignalService.FindAppOrCreate(pushPacket.Options.SiteURL)
	if err != nil {
		return nil, err
	}
	// Get the App ID for further API calls
	notificationRequest := &onesignal.NotificationRequest{}
	notificationRequest.AppID = foundApp.ID
	notificationRequest.Contents = map[string]string{"en": pushPacket.Options.Text}
	notificationRequest.Headings = map[string]string{"en": pushPacket.Options.Title}
	// TODO: figure out which one to use
	notificationRequest.IsAnyWeb = true
	notificationRequest.IsIOS = true
	notificationRequest.IsAndroid = true
	notificationRequest.IncludedSegments = []string{"Active Users", "Inactive Users"}
	// REST API key is used on per-app basis
	app.Services.OneSignalService.SetAppRestAuthKey(foundApp.BasicAuthKey)
	notificationResponse, err := app.Services.OneSignalService.SendNotification(notificationRequest)

	if err != nil {
		return nil, err
	} else {
		return notificationResponse, nil
	}
}
