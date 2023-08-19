package entity

type SendGridBaseInfo struct {
	FromName    string
	FromAddress string
	ToName      string
	ToAddress   string
	Subject     string
}

// 通常のメール送信時用
type SendInfo struct {
	BaseInfo         SendGridBaseInfo
	PlainTextContent string
	HtmlContent      string
}

// Dynamic Templateでのメール送信用
type SendInfoWithTemplate struct {
	BaseInfo   SendGridBaseInfo
	TemplateID string
}
