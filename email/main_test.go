package email

import (
	"testing"
)

func TestContainer(t *testing.T) {
	conf := Configure{
		Host:     "smtp.bizmail.yahoo.com",
		Port:     587,
		UserName: "test",
		Password: "test",
		From:     "test",
	}
	to := Model{
		To:      "test",
		ToName:  "Tho",
		Subject: "test",
		Body:    "test cc",
		CC:      "test",
		BCC:     "test",
	}
	r := New(conf)
	err := r.SendEmail(&to)
	if err != nil {
		t.Fatal("send email error")
	}
}
