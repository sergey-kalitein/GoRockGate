package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload" // for the main ".env" file
	"net/http"
	"os"
	"rockgate/app"
	"rockgate/controllers"
)

func main() {
	// Configuring
	app.Configure()
	// Init all services
	app.InitServices()
	// Setting up the routing
	setupRouter()
	// Launching the server
	serve()
}

func setupRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/push/{service:[a-z0-9]+}/send", controllers.SendPushNotification)
	router.HandleFunc("/apps/find-or-create/{domain:[^/]+}", controllers.FindOrCreateApplication)
	router.HandleFunc("/apps/find/{domain:[^/]+}", controllers.FindApplication)
	router.HandleFunc("/apps/list", controllers.ListApplications)
	router.HandleFunc("/status", HandlerServerStatus)
	http.Handle("/", router)
}

func serve() {
	fmt.Println("Push Gateway is listening...")
	gatewayAddress := os.Getenv("SERVICE_ADDRESS")
	if gatewayAddress == "" {
		gatewayAddress = ":8181"
	}
	http.ListenAndServe(gatewayAddress, nil)
}
