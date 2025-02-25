package email

import (
	"log"
	"os"
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) {
    if to == "" {
        log.Println("❌ Skipping email: No recipient address provided")
        return
    }

    smtpEmail := os.Getenv("SMTP_EMAIL")
    smtpHost := os.Getenv("SMTP_HOST")
    smtpPassword := os.Getenv("SMTP_PASSWORD")

    if smtpEmail == "" || smtpHost == "" || smtpPassword == "" {
        log.Println("❌ Missing SMTP credentials")
        return
    }

    m := gomail.NewMessage()
    m.SetHeader("From", smtpEmail)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)

    d := gomail.NewDialer(smtpHost, 587, smtpEmail, smtpPassword)

    if err := d.DialAndSend(m); err != nil {
        log.Println("❌ Failed to send email:", err)
    } else {
        log.Println("📩 Email sent successfully to", to)
    }
}