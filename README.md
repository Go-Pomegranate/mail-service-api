# mail-service-api
This is mail service which combine two 3th party APIs for sending mails via net - sendGrid and mailGun. Main purpose of that service is to create services from those API's and jump between them in case of the failover.

# Getting Started

What hasn't been done?

- not sharing app on Heroku etc. no dockerization ... -> sorry about that , unfortunetly clock was ticking for me.
It means that to test app, You need your own apiKey and domain for those 3th party API's -> whispering: or You can just contact me - GitGuardian doesn't hear... :) 

- front-end is reduced to just silly "Welcome" -> all prints logs are in the console

- no Hystrix-Go circuit breaker made good way -> no service running on the background  just implementation of the service inside project. It was my idea to how to handle failover.


- no data from MongoDB, just little fake repo imitating that it is connecting

What has been done?

- mail service which use 2 API's -> sendGrid and mailGun. It recovers when one API is down via circuit breaker pattern made "by hand" and another one with just implementation of  Hystrix-Go , but without service in the background. So Hystrix-Go was blocking goroutines and I was forced to just leave an implementation.

- standard logger to log all web requests on the console

- fake repository to MongoDB -> why MongoDB? It is simple for that scale of project, gives nice flexibility with JSON when requirements of project will change and it is very near to my OOP level skills with querying the database.

- configuration struct which reading from config.json file ( you can pass your apiKey and domainName there to make it run)

- tests for mail service, test for failover for that service and test for handler of "/index" GET request

- mail service interface -> it's something which i'm not still very familiar with, so i limited my interface Golang experience just to write simple interface and one method that services implements ( SendMail() )

- simple error handling -> propagation to upstream and logging

# Prerequisites

I just want to mention here:

**that I am Golang beginner and came from Java World, where some place are different and code smells in Go are not code smells in Java. 

**Pretty much time I lost on structurizing the project directories to be not ashamed of it as much.


# Installing

To get all necessary libraries just do those command on the console:

go get github.com/tkanos/gonfig 

go get gopkg.in/mgo.v2

go get github.com/gorilla/mux

go get github.com/afex/hystrix-go/hystrix

go get "github.com/mailgun/mailgun-go"
	
go get "github.com/sendgrid/sendgrid-go"


# Built With
IDEA Goland 2.3 -> golang IDE
mgo.v2 -> mongoDB driver
tkanos/gonfig -> JSON based configuration 
gorilla/mux -> package for simpler web handlers and routers
hystrix-go/hystrix -> circuit breaker pattern package
mailgun-go -> API mailGun for sending mails
sendgrid-go -> API sendGrid for sending mails

# Author
Eryk Panter

# License
This project is licensed under the MIT License - see the LICENSE.md file for details
