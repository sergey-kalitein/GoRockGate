package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"os"
	"rockgate/app"
	"rockgate/controllers"
)

func main() {
	// Configuring
	configure()
	// Init all services
	app.InitServices()
	// Setting up the routing
	setupRouter()
	// Launching the server
	serve()
}

func serve() {
	fmt.Println("Push Gateway is listening...")
	gatewayAddress := os.Getenv("SERVICE_ADDRESS")
	if gatewayAddress == "" {
		gatewayAddress = ":8181"
	}
	http.ListenAndServe(gatewayAddress, nil)
}

func setupRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/push/{service:[a-z0-9]+}/send", controllers.HandlerPushNotifications)
	router.HandleFunc("/apps/find-or-create/{domain:[^/]+}", controllers.FindOrCreateApplication)
	router.HandleFunc("/apps/list", controllers.ListApplications)
	router.HandleFunc("/status", HandlerServerStatus)
	http.Handle("/", router)
}

func configure() {
	configWarnings, err := app.LoadConfiguration()
	if err != nil {
		app.Fatal(err.Error())
	}
	if len(configWarnings) > 0 {
		for _, warningText := range configWarnings {
			color.New(color.FgBlack, color.BgYellow).Printf("[CONFIG WARNING] %s", warningText)
			fmt.Println()
		}
	}
}
