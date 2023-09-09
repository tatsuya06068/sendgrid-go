package driver

import (
	"context"
	"example/sendgrid/entity"
	"fmt"
	"log"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type IClient interface {
	SendWithContext(ctx context.Context, email *mail.SGMailV3) (*rest.Response, error)
}

type SendGridDriver struct {
	client IClient
}

func NewSendGridDriver(apiKey string) SendGridDriver {
	client := sendgrid.NewSendClient(apiKey)
	return SendGridDriver{
		client: client,
	}
}

func (sgd *SendGridDriver) SendMail(ctx context.Context, sendInfo entity.SendInfo) error {

	from := mail.NewEmail(sendInfo.BaseInfo.FromName, sendInfo.BaseInfo.FromAddress)
	subject := sendInfo.BaseInfo.Subject
	to := mail.NewEmail(sendInfo.BaseInfo.ToAddress, sendInfo.BaseInfo.FromAddress)
	plainTextContent := sendInfo.PlainTextContent
	htmlContent := sendInfo.HtmlContent

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	response, err := sgd.client.SendWithContext(ctx, message)

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

func (sgd *SendGridDriver) SendMailWithTemplate(ctx context.Context, sendInfo entity.SendInfoWithTemplate, tempValues ...map[string]any) error {

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

	response, err := sgd.client.SendWithContext(ctx, message)

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
