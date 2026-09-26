package users

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type GomailUtil struct{}

func CreateGomailUtil() *GomailUtil {
	return &GomailUtil{}
}

func (gu GomailUtil) SendEmail(email SendEmailParams) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SENDER"))
	m.SetHeader("To", email.Receiver)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.EmailBody)

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}

	d := gomail.NewDialer(os.Getenv("SMTP_SERVER"), port, os.Getenv("SENDER"), os.Getenv("SENDER_PASSWORD"))
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
