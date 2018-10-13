// package models represents structures e.g structs which describes entities used in whole mail service
package models

// Mail is simple entity that represents mails transported by mail service.
// "Sender" attribute is mail address of the sender e.g John.Smith@domain.com.
// "Receiver" attribute is mail address of the receiver e.g John.Smith@domain.com.
// "Subject" is mail topic
// "Content" content of mail
type Mail struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Subject  string `json:"subject"`
	Content  string `json:"content"`
}

// this slice repesents group of user's mails
type Mails []Mail
