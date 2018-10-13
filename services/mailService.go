// package services implements whole app's services and their methods. Consists circuit breaker pattern service.
package services

import (
	"github.com/Go_Pomegranate/config"
	"github.com/Go_Pomegranate/models"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/mailgun/mailgun-go"
	"github.com/pkg/errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"strconv"
)

// init of configuration file
var configuration = config.InitConfig()

// custom error
var unableToSendMailViaSendGrid = errors.New("Unable to send the email via SendGrid")

// represents standard mail service without circuit breaker pattern
type MailService struct {
}

// method to send mails failover service
// require Mail object as argument
// returns status code of http request e.g 200 OK and error object if exception occured
func (service *MailService) SendMail(mailMessage models.Mail) (string, error) {

	result, err := sendGunMail(mailMessage)

	if err != nil {
		log.Println("sendGunMail Service shutted down.. \n" + err.Error())
		log.Println("Starting new backup service sendGridService...")
		result, err := sendGridMail(mailMessage)

		if err != nil {
			log.Println("Base and backup service are off. We are doomed.")
			return "", err
		}

		if result != "200" && result != "201" && result != "202" {
			log.Println(result)
			log.Println("Base and backup service are off. We are doomed.")
			return "", err
		}

		log.Println("Backup service successfully ON.")

		return result, nil

	}

	return result, nil
}

// function to send mails via GunMail third party Vendor service
// require Mail object as argument
// returns status code of http request e.g 200 OK and error object if exception occured
func sendGunMail(mailMessage models.Mail) (string, error) {

	log.Println("Sending via gunMail..")

	mailGun := mailgun.NewMailgun(configuration.MailGunDomainURL, configuration.MailGunApiKey)
	mailDetails := mailGun.NewMessage(
		mailMessage.Sender,
		mailMessage.Subject,
		mailMessage.Content,
		mailMessage.Receiver,
	)

	result, id, err := mailGun.Send(mailDetails)
	if err != nil {
		return id, err
	}

	return result, err
}

// function to send mails via SendGrid third party Vendor service
// require Mail object as argument
// returns status code of http request e.g 200 OK and error object if exception occure
func sendGridMail(mailMessage models.Mail) (string, error) {

	log.Println("Sending via sendGrid..")

	from := mail.NewEmail(mailMessage.Sender, mailMessage.Sender)
	subject := mailMessage.Content
	to := mail.NewEmail(mailMessage.Receiver, mailMessage.Receiver)
	plainTextContent := mailMessage.Content
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(configuration.SendGridApiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return unableToSendMailViaSendGrid.Error(), err
	} else {
		return strconv.Itoa(response.StatusCode), nil
	}
	return "Something unpredicted happend", err
}

// this struct represents circuit breaker pattern service for sending mails
// there are 2 available third party services now: SendGrid and GunMail
// if one of those services failed then backup service will be started
type MailServiceWithCircuitBreaker struct {
}

// method to send mails with circuit breaker failover service
// require Mail object as argument
// returns status code of http request e.g 200 OK and error object if exception occured
func (service *MailServiceWithCircuitBreaker) SendMail(mailMessage models.Mail) (string, error) {

	output := make(chan string, 1)
	hystrix.ConfigureCommand("send_mail", hystrix.CommandConfig{Timeout: 1000})
	chanerrors := hystrix.Go("send_mail", func() error {
		_, err := sendGunMail(mailMessage)

		if err != nil {
			log.Println("sendGunMail Service broke circuit... \n" + err.Error())
			return err
		}

		return nil

	}, func(err error) error {
		_, err = sendGridMail(mailMessage)
		if err != nil {
			return err
		}
		return nil
	})

	select {
	case out := <-output:
		return out, nil

	case err := <-chanerrors:
		return unableToSendMailViaSendGrid.Error(), err
	}

}
