package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/tbalthazar/onesignal-go"
	"log"
	"net/http"
)

var oneSignalService *OneSignalService

var oneSignalApps *[]onesignal.App

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
	router.HandleFunc("/push/{service:[a-z0-9]+}/send", HandlerPushNotifications)
	router.HandleFunc("/status", HandlerServerStatus)
	http.Handle("/", router)
}

func loadOneSignalApps() {
	oneSignalService = NewOneSignalService(Config())
	_, err := oneSignalService.LoadAppsList()
	if err != nil {
		fatal("[OneSignal::ListApps]" + err.Error())
	}
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
