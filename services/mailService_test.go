package services

import (
	"github.com/Go_Pomegranate/models"
	"log"
	"testing"
)

func TestMailService(t *testing.T) {

	mail := models.Mail{Sender: "postmaster@sandboxdf5be3d8493a4c6e89ed2cb5ed724f65.mailgun.org", Receiver: "eryk.panter@gmail.com", Content: "Hello from the other side!"}

	mailService := MailService{}
	result, err := mailService.SendMail(mail)
	if err != nil {
		log.Fatal(err)
	}
	if result == "Queued. Thank you." {
		log.Println("Test passed successfully")
	}

}

func TestMailServiceWithCircuitBreaker(t *testing.T) {

	mail := models.Mail{Sender: "postmaster@sandboxdf5be3d8493a4c6e89ed2cb5ed724f65.mailgun.orgfdgdfgdfgdg", Receiver: "eryk.panter@gmail.com", Content: "Hello from the other side!"}

	mailService := MailService{}
	result, err := mailService.SendMail(mail)
	if err != nil {
		log.Fatal(err)
	}
	if result == "Queued. Thank you." {
		log.Println("Test passed successfully")
	}

	if result == "202" {
		log.Println("Test passed successfully with circuit breaker backup service sendGrid.")
	}

}
