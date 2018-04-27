package main

import (
	"fmt"
	"log"
	"net/http"
	"fibonacci/common"
	"fibonacci/routers"
)

func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	// Create the Server
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Printf("Listening... on port: %s", server.Addr)
	// Running the HTTP Server
	go server.ListenAndServe()
	go http.ListenAndServeTLS(":4443", "cert.pem", "key.pem", nil)

	fmt.Scanln()
}
