package handler

import (
	"log"

	"github.com/clubo-app/protobuf/events"
	mail "github.com/xhit/go-simple-mail/v2"
)

const htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello Gophers!</title>
	</head>
	<body>
		<p>This is the <b>Go gopher</b>.</p>
		<p><img src="cid:Gopher.png" alt="Go gopher" /></p>
		<p>Image created by Renee French</p>
	</body>
</html>`

func (s *server) Registered(m *events.Registered) {
	log.Printf("%v", m)
	email := mail.NewMSG()
	email.SetFrom(s.config.EMAIL_FROM).
		AddTo(m.Email).
		SetSubject("Verify your Email")

	email.SetBodyData(mail.TextHTML, []byte(htmlBody))

	// always check error after send
	if email.Error != nil {
		log.Println(email.Error)
	}

	// Call Send and pass the client
	err := email.Send(s.mail)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}

}
