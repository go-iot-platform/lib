package email

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

/*
Model define model
*/
type Model struct {
	To      string
	ToName  string
	Subject string
	Body    string
	CC      string
	BCC     string
}

// Configure email
type Configure struct {
	Host     string
	Port     int
	UserName string
	Password string
	From     string
}

// Receiver resprensent
type Receiver struct {
	Sender gomail.Sender
	Dialer *gomail.Dialer
	From   string
}

/*
New init parameter
*/
func New(conf Configure) *Receiver {

	d := gomail.NewDialer(conf.Host, conf.Port, conf.UserName, conf.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	res := &Receiver{
		Dialer: d,
		From:   conf.From,
	}
	return res
}

/*
SendEmail send email
*/
func (r *Receiver) SendEmail(p *Model) error {
	m := gomail.NewMessage()

	m.SetAddressHeader("From", r.From, "Onsky Inc - Cloud Support")
	m.SetAddressHeader("To", p.To, p.ToName)
	if p.CC != "" {
		m.SetAddressHeader("CC", p.CC, "")
	}
	m.SetHeader("Subject", p.Subject)
	m.SetBody("text/html", p.Body)

	s, err := r.Dialer.Dial()
	if err != nil {
		return err
	}

	if err := gomail.Send(s, m); err != nil {
		return err
	}
	m.Reset()
	return nil
}
