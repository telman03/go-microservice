package email

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

// SendEmail sends an email notification
func SendEmail(to string, subject string, body string) {
    smtpEmail := os.Getenv("SMTP_EMAIL")
    smtpHost := os.Getenv("SMTP_HOST")
    smtpPassword := os.Getenv("SMTP_PASSWORD")

    log.Println("SMTP_EMAIL:", smtpEmail)
    log.Println("SMTP_HOST:", smtpHost)

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
        log.Println("📩 Email sent successfully!")
    }
}

// HandleMessage processes Kafka messages and sends email
func HandleMessage(message []byte) {
	log.Println("Processing message:", string(message))

	// Send email
	SendEmail("telmangadimov7@gmail.com", "New Order Received", "Your order has been placed successfully!")
}