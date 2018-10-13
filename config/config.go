// package config store all necessary const variables like
// DB connection info or vendors mail service API keys in entities like struct
package config

import (
	"github.com/tkanos/gonfig"
	"log"
)

// Configuration stores configuration variables
type Configuration struct {
	ServerURL        string
	MongoDBName      string
	SendGridApiKey   string
	MailGunDomainURL string
	MailGunApiKey    string
}

// function reads from config.json file and initialize all attributes of Configuration struct
func InitConfig() *Configuration {
	config := Configuration{}
	err := gonfig.GetConf("C:\\Users\\Eryk\\Documents\\goworkspace\\src\\github.com\\Go_Pomegranate\\config.json", &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}
