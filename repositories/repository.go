// package repositories connects to datasources and defines all DB requests methods on that datasources
package repositories

import (
	"github.com/Go_Pomegranate/config"
	"github.com/Go_Pomegranate/models"
	"gopkg.in/mgo.v2"
	"log"
)

// MailRepository is repository handling all mail service queries
type MailRepository struct {
}

var configuration = config.InitConfig()

// defines DB url
var SERVER = configuration.ServerURL

// defines DB name
var DBNAME = configuration.MongoDBName

// defines MongoDB's document name which holds all mail records
const DOCNAME = "mails"

// Function GetMails() returns all mails from DB document, limit to 50
func (r MailRepository) GetMails() models.Mails {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Println("Failed to establish connection to MongoDB:", err)
		return models.Mails{}
	}

	defer session.Close()

	conn := session.DB(DBNAME).C(DOCNAME)
	results := models.Mails{}
	// limit Mails returned to 50 for performance
	if err := conn.Find(nil).Limit(50).All(&results); err != nil {
		log.Println("Failed to write results:", err)
		return models.Mails{}
	}
	return results
}
