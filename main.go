// main package which is responsible for starting Web Server and initiliazing services
package main

import (
	"github.com/Go_Pomegranate/routers"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func main() {

	//mail := models.Mail{Sender:"postmaster@sandboxdf5be3d8493a4c6e89ed2cb5ed724f65.mailgun.org", Receiver: "eryk.panter@gmail.com", Content:"Hello from the other side!"}
	//
	////mailService := services.MailServiceWithCircuitBreaker{}
	//
	//mailService := services.MailService{}
	//
	//result, err := mailService.SendMail(mail)
	//
	//if err != nil {
	//	log.Printf("%+v\n", err)
	//}
	//
	//log.Println(result)

	router := routers.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":8000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))

}
