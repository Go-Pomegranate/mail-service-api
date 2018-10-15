// main package which is responsible for starting Web Server and initiliazing services
package main

import (
	"github.com/Go_Pomegranate/routers"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func main() {

	router := routers.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":8000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))

}
