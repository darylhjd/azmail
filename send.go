package azmail

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type mailMessage struct {
	Attachments                    []MailAttachment `json:"attachments"`
	Content                        MailContent      `json:"content"`
	Recipients                     MailRecipients   `json:"recipients"`
	ReplyTo                        []MailAddress    `json:"replyTo"`
	SenderAddr                     string           `json:"senderAddress"`
	UserEngagementTrackingDisabled bool             `json:"userEngagementTrackingDisabled"`
}

func (c *Client) newMailMessage(mail Mail) mailMessage {
	return mailMessage{
		nil,
		mail.Content,
		mail.Recipients,
		nil,
		c.senderAddr,
		true,
	}
}

// SendMails sends multiple mails. If any errors are encountered, the error is saved and later returned.
// Encountering errors does not stop later emails from being sent.
func (c *Client) SendMails(mails ...*Mail) error {
	var errs []error

	for _, mail := range mails {
		msg := c.newMailMessage(*mail)
		if err := c.sendMessage(msg); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

type errorResponse struct {
	Error struct {
		AdditionalInfo []struct {
			Info any    `json:"info"`
			Type string `json:"type"`
		} `json:"additionalInfo"`
		Code    string          `json:"code"`
		Details []errorResponse `json:"details"`
		Message string          `json:"message"`
		Target  string          `json:"target"`
	} `json:"error"`
}

func (c *Client) sendMessage(msg mailMessage) error {
	req, err := c.generateSignedMessageRequest(msg)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusAccepted {
		return nil
	}

	var (
		b       bytes.Buffer
		errResp errorResponse
	)
	if _, err = b.ReadFrom(resp.Body); err != nil {
		return err
	}

	if err = json.Unmarshal(b.Bytes(), &errResp); err != nil {
		return err
	}

	return errors.New(errResp.Error.Message)
}
