package driver

import (
	"example/sendgrid/entity"
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridDriver struct {
	ApiKey string
}

func NewSendGridDriver(apiKey string) SendGridDriver {
	return SendGridDriver{
		ApiKey: apiKey,
	}
}

func (sgd *SendGridDriver) SendMail(sendInfo entity.SendInfo) error {

	from := mail.NewEmail(sendInfo.BaseInfo.FromName, sendInfo.BaseInfo.FromAddress)
	subject := sendInfo.BaseInfo.Subject
	to := mail.NewEmail(sendInfo.BaseInfo.ToAddress, sendInfo.BaseInfo.FromAddress)
	plainTextContent := sendInfo.PlainTextContent
	htmlContent := sendInfo.HtmlContent

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(sgd.ApiKey)
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return nil

}

func (sgd *SendGridDriver) SendMailWithTemplate(sendInfo entity.SendInfoWithTemplate, tempValues ...map[string]any) error {

	from := mail.NewEmail(sendInfo.BaseInfo.FromName, sendInfo.BaseInfo.FromAddress)
	subject := sendInfo.BaseInfo.Subject
	to := mail.NewEmail(sendInfo.BaseInfo.ToAddress, sendInfo.BaseInfo.FromAddress)

	message := mail.NewV3MailInit(from, subject, to)
	message.SetTemplateID(sendInfo.TemplateID)

	if len(tempValues) != 0 {
		p := mail.NewPersonalization()

		for _, valueMap := range tempValues {
			for key, value := range valueMap {
				p.SetDynamicTemplateData(key, value)
			}
		}

		message.AddPersonalizations(p)
	}

	client := sendgrid.NewSendClient(sgd.ApiKey)
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return nil
}
