package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/tbalthazar/onesignal-go"
	"log"
	"net/http"
	"rockgate/handlers"
)

var oneSignalService *OneSignalService

var oneSignalApps *[]onesignal.App

func HandlerServerStatus(responseWriter http.ResponseWriter, request *http.Request) {
	// TODO: Check whether we are Ok or we can not create new Apps due to the lack of configuration data
}

func processOneSignal() {
	client := onesignal.NewClient(nil)
	// TODO: get the user key from the settings
	client.UserKey = "OneSignalUserKey"
	// Find an App
	// Get an App Key by the domain name
	// Config all certificates/keys/tokens from the settings
	// Create an App if there is no App

	// When we have a specific app:
	client.AppKey = "YourOneSignalAppKey"
}

func main() {
	// Configuring
	configure()
	// Load Applications
	loadOneSignalApps()
	// Setting up the routing
	setupRouter()
	// Launching the server
	serve()
}

func serve() {
	fmt.Println("Push Gateway is listening...")
	http.ListenAndServe(":8181", nil)
}

func setupRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/push/{service:[a-z0-9]+}/send", handlers.HandlerPushNotifications)
	router.HandleFunc("/status", HandlerServerStatus)
	http.Handle("/", router)
}

func loadOneSignalApps() {
	oneSignalService = NewOneSignalService(Config())
	apps, _, err := oneSignalService.ListApps()
	if err != nil {
		fatal("[OneSignal::ListApps]" + err.Error())
	}
	oneSignalApps = &apps
}

func configure() {
	configWarnings, err := LoadConfiguration()
	if err != nil {
		fatal(err.Error())
	}
	if len(configWarnings) > 0 {
		for _, warningText := range configWarnings {
			color.New(color.FgBlack, color.BgYellow).Printf("[CONFIG WARNING] %s", warningText)
			fmt.Println()
		}
	}
}

func fatal(errorText string) {
	color.New()
	log.Fatal(color.New(color.BgRed, color.FgHiYellow).Sprintf("[FATAL ERROR] %s", errorText))
}
