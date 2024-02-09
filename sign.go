package azmail

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func (c *Client) generateSignedMessageRequest(msg mailMessage) (*http.Request, error) {
	// https://learn.microsoft.com/en-us/rest/api/communication/dataplane/email/send?view=rest-communication-dataplane-2023-03-31&viewFallbackFrom=rest-communication-dataplane-2023-04-01-preview&tabs=HTTP
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// https://learn.microsoft.com/en-us/rest/api/communication/authentication
	pathAndQuery := fmt.Sprintf("%s?%s", c.u.Path, c.u.Query().Encode())

	timestamp := strings.ReplaceAll(time.Now().UTC().Format(time.RFC1123), "UTC", "GMT")

	hash := sha256.Sum256(body)
	hashB64 := base64.StdEncoding.EncodeToString(hash[:])

	stringToSign := fmt.Sprintf(
		"%s\n%s\n%s;%s;%s",
		http.MethodPost, pathAndQuery, timestamp, c.u.Host, hashB64,
	)

	hm := hmac.New(sha256.New, c.accessKey)
	if _, err = hm.Write([]byte(stringToSign)); err != nil {
		return nil, err
	}

	signature := base64.StdEncoding.EncodeToString(hm.Sum(nil))
	authorization := fmt.Sprintf("HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s", signature)

	req.Header.Set("Content-Type", "application/json")
	req.Header["x-ms-date"] = []string{timestamp}
	req.Header["x-ms-content-sha256"] = []string{hashB64}
	req.Header["host"] = []string{pathAndQuery}
	req.Header.Set("Authorization", authorization)

	return req, nil
}
