package models

type Email struct {
	To       string `json:"to"`
	Subject  string `json:"subject"`
	BodyText string `json:"body_text"`
	BodyHTML string `json:"body_html"`
}

type MagicLinkEmailData struct {
	MagicLink string
	Name      string
	Year      int
}
