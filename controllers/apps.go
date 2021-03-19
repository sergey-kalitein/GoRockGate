package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"rockgate/app"
)

func ListApplications(responseWriter http.ResponseWriter, request *http.Request) {
	apps, err := app.Services.OneSignalService.GetApps()
	if err != nil {
		app.SendOutError(responseWriter, err.Error(), http.StatusBadRequest)
	} else {
		app.SendOutJSON(responseWriter, *apps, 200)
	}
}

func FindApplication(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	originDomain := vars["domain"]
	oneSignalApp, isFound := app.Services.OneSignalService.FindApp(originDomain)
	if isFound == true {
		app.SendOutJSON(responseWriter, oneSignalApp, 200)
	} else {
		app.SendOutError(responseWriter, "application not found", http.StatusNotFound)
	}
}

func FindOrCreateApplication(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	originDomain := vars["domain"]
	oneSignalApp, err := app.Services.OneSignalService.FindAppOrCreate(originDomain)
	if err != nil {
		app.SendOutError(responseWriter, err.Error(), http.StatusBadRequest)
	} else {
		app.SendOutJSON(responseWriter, oneSignalApp, 200)
	}
}
