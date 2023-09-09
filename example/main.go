package main

import (
	"context"
	"example/sendgrid/driver"
	"example/sendgrid/entity"
	"log"
	"os"
)

func main() {
	dri := driver.NewSendGridDriver(os.Getenv("SENDGRID_API_KEY"))
	sendMail(dri)
}

func sendMail(driver driver.SendGridDriver) {
	baseInfo := entity.SendGridBaseInfo{
		FromName:    "",
		FromAddress: "",
		ToName:      "",
		ToAddress:   "",
		Subject:     "",
	}

	info := entity.SendInfo{
		BaseInfo:         baseInfo,
		PlainTextContent: "test",
		HtmlContent:      "<strong>test</strong>",
	}

	err := driver.SendMail(context.Background(), info)

	if err != nil {
		log.Fatal(err)
	}
}

func SendMailWithTemplate(driver driver.SendGridDriver) {
	baseInfo := entity.SendGridBaseInfo{
		FromName:    "",
		FromAddress: "",
		ToName:      "",
		ToAddress:   "",
		Subject:     "",
	}

	info := entity.SendInfoWithTemplate{
		BaseInfo:   baseInfo,
		TemplateID: os.Getenv("TEMPLATE_ID"),
	}

	err := driver.SendMailWithTemplate(context.Background(), info)

	if err != nil {
		log.Fatal(err)
	}
}
