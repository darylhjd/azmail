package azmail

import (
	"encoding/base64"
	"net/url"
)

const (
	apiPath    = "/emails:send"
	apiVersion = "2023-03-31"
)

// Client for the  Azure Email Communication Service.
type Client struct {
	u         *url.URL
	accessKey []byte

	senderAddr string
}

// NewClient creates a new client for Azure Email Communication Service.
// Use the provided endpoint, access key, and email address from your communication service.
func NewClient(endpoint, accessKey, senderAddr string) (*Client, error) {
	rawKey, err := base64.StdEncoding.DecodeString(accessKey)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Set("api-version", apiVersion)
	u.RawQuery = v.Encode()
	u.Path = apiPath

	return &Client{
		u:          u,
		accessKey:  rawKey,
		senderAddr: senderAddr,
	}, nil
}
