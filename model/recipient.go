package model

type Recipient struct {
	Type  string `json:"type"`
	Email string `json:email"`
	Name  string `json:name`
}