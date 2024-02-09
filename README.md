# azmail

[![Go Reference](https://pkg.go.dev/badge/github.com/darylhjd/azmail.svg)](https://pkg.go.dev/github.com/darylhjd/azmail)
[![Go Report](https://goreportcard.com/badge/github.com/darylhjd/azmail?style=flat-square)](https://goreportcard.com/report/github.com/darylhjd/azmail)

Go API for sending emails through Microsoft Azure's Email Communication Service.

## Installation

```cmd
$ go get -u github.com/darylhjd/azmail
```

## Example Usage

```go
package main

import (
	"log"

	"github.com/darylhjd/azmail"
)

func main() {
	client, _ := azmail.NewClient("ENDPOINT", "ACCESS_KEY", "SENDER_ADDRESS")

	// Create mails that you want to send.
	mail1 := azmail.NewMail()
	mail1.Recipients = ...
	mail1.Content = ...
	mail1.Attachments = ...
	mail2 := azmail.NewMail()
	...

	// Send your mails.
	errs := client.SendMails(mail1, mail2)
	log.Println(errs)
}
```

## Current Version and Documentation

This wrapper implements version `2023-03-31` of the Email API.

More information on the API can be
found [here](https://learn.microsoft.com/en-us/rest/api/communication/dataplane/email/send?view=rest-communication-dataplane-2023-03-31&viewFallbackFrom=rest-communication-dataplane-2023-10-01&tabs=HTTP).