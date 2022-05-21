package model

type Payload struct {
	SenderMail  string      `json:"sender_mail"`
	SenderName  string      `json:"sender_name"`
	Recipients  []Recipient `json:"recipients"`
	ReplyEmail  string      `json:"reply_email"`
	Subject     string      `json:"subject"`
	BodyHTML    string      `json:"body_html"`
	BodyText    string      `json:"body_text"`
	Attachments []string    `json:"attachments"`
}