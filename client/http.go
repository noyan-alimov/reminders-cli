package client

import (
	"net/http"
	"time"
)

// NewHTTPClient - constructor for the client
func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		BackendURI: uri,
		client:     &http.Client{},
	}
}

// HTTPClient - client
type HTTPClient struct {
	client     *http.Client
	BackendURI string
}

// Create - calls Create api endpoint
func (c HTTPClient) Create(title, message string, duration time.Duration) ([]byte, error) {
	res := []byte(`response for Create reminder`)
	return res, nil
}

// Edit - calls Edit api endpoint
func (c HTTPClient) Edit(id, title, message string, duration time.Duration) ([]byte, error) {
	res := []byte(`response for Edit reminder`)
	return res, nil
}

// Fetch - calls Fetch api endpoint
func (c HTTPClient) Fetch(ids []string) ([]byte, error) {
	res := []byte(`response for Fetch reminder`)
	return res, nil
}

// Delete - calls Delete api endpoint
func (c HTTPClient) Delete(ids []string) error {
	return nil
}

// Healthy - calls Healthy api endpoint
func (c HTTPClient) Healthy(host string) bool {
	return true
}
