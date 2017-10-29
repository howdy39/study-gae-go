package mail

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
	"net/http"
)

const subject = "Test Send"
const mailBody = `
Thank you!

This is demo of Sending Email.
`

func init() {
	http.HandleFunc("/sendEmail", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	sender := "admin@howdylikes.jp"
	toAddress := []string{"admin@howdylikes.jp"}

	msg := &mail.Message{
		Sender:  sender,
		To:      toAddress,
		Subject: subject,
		Body:    mailBody,
	}

	if err := mail.Send(ctx, msg); err != nil {
		fmt.Print(w, "エラーが発生しました")
	}
}
