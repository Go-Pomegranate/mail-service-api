// package controllers handle all requests from Web
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Go_Pomegranate/repositories"
	"log"
	"net/http"
)

// Controller is basic struct which holds all repositories for handlers
type Controller struct {
	repositories.MailRepository
}

// function Index handles request of getting index Web content to the page
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	mails := c.MailRepository.GetMails()
	log.Println(mails)
	data, _ := json.Marshal(mails)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

	fmt.Fprint(w, "Welcome! \n")
	return
}
