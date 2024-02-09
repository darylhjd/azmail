package azmail

type Mail struct {
	Recipients  MailRecipients   `json:"recipients"`
	Content     MailContent      `json:"content"`
	Attachments []MailAttachment `json:"attachments"`
}

type MailRecipients struct {
	To  []MailAddress `json:"to"`
	Cc  []MailAddress `json:"cc"`
	Bcc []MailAddress `json:"bcc"`
}

type MailAddress struct {
	Address     string `json:"address"`
	DisplayName string `json:"displayName"`
}

type MailContent struct {
	Subject   string `json:"subject"`
	PlainText string `json:"plainText"`
	Html      string `json:"html"`
}

type MailAttachment struct {
	Name          string `json:"name"`
	Base64Content string `json:"contentInBase64"`
	ContentType   string `json:"contentType"`
}

// NewMail is a convenience function for creating a new Mail and returning the pointer to it.
func NewMail() *Mail {
	return &Mail{}
}
