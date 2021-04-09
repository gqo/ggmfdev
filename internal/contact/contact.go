package contact

import (
	"log"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	Address string `json:"email_address"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

var (
	contactEmail   string
	sendgridAPIKey string
)

func SendEmail(email Email) error {
	from := mail.NewEmail(email.Name, email.Address)
	to := mail.NewEmail("Graeme Ferguson", contactEmail)

	message := mail.NewSingleEmail(from, email.Subject, to, email.Message, email.Message)

	client := sendgrid.NewSendClient(sendgridAPIKey)

	response, err := client.Send(message)
	switch response.StatusCode {
	case http.StatusAccepted, http.StatusOK:
		return nil
	default:
		return err
	}
}

func init() {
	log.Println("contact: init(): Starting environement variable fetch.")

	if value, ok := os.LookupEnv("GGMFDEV_CONTACT_EMAIL"); ok {
		contactEmail = value
	} else {
		log.Fatalln("contact: init(): Failed to fetch GGMFDEV_CONTACT_EMAIL from environment.")
	}

	if value, ok := os.LookupEnv("GGMFDEV_SENDGRID_API_KEY"); ok {
		sendgridAPIKey = value
	} else {
		log.Fatalln("contact: init(): Failed to fetch GGMFDEV_SENDGRID_API_KEY from environment.")
	}

	log.Println("contact: init(): Finished.")
}
