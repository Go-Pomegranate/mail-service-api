// package interfaces provides interface abstractions for services
package interfaces

import "github.com/Go_Pomegranate/models"

// IMailService is interface of mail service
type IMailService interface {
	SendMail(mail models.Mail) (string, error)
}
