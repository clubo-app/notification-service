package mail

import (
	"time"

	"github.com/clubo-app/notification-service/config"
	mail "github.com/xhit/go-simple-mail/v2"
)

func Connect(c config.Config) (*mail.SMTPClient, error) {
	server := mail.NewSMTPClient()

	// SMTP Server
	if c.SMTP_HOST != "" {
		server.Host = c.SMTP_HOST
	}
	if c.SMTP_PORT != 0 {
		server.Port = c.SMTP_PORT
	}
	if c.SMTP_USERNAME != "" {
		server.Username = c.SMTP_USERNAME
	}
	if c.SMTP_PW != "" {
		server.Password = c.SMTP_PW
	}

	// Since v2.3.0 you can specified authentication type:
	// - PLAIN (default)
	// - LOGIN
	// - CRAM-MD5
	// - None
	// server.Authentication = mail.AuthPlain

	// Variable to keep alive connection
	server.KeepAlive = false

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	server.SendTimeout = 10 * time.Second

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	// server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// SMTP client
	return server.Connect()
}
