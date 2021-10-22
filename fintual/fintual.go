// Package fintual provides utilties for interfacing
// with the Fintual API.
package fintual

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Version is the version of this library.
const Version = "1.0.0"

const (
	baseURL = "https://fintual.cl/api"
)

type Client struct {
	http        *http.Client // HTTP client used to communicate with the API.
	baseURL     *url.URL     // Base URL for API requests
	accessToken string       // Access token used for methods which require authentication
}

// NewClient returns a new Fintual API client.
// To use API methods which require authentication, you must first call
// the provided Authenticate method with valid credentials.
func NewClient() *Client {
	baseURL, _ := url.Parse(baseURL)
	return &Client{http: &http.Client{Timeout: time.Minute}, baseURL: baseURL}
}

// Error represents an error returned by the Fintual API.
type Error struct {
	Code    int    `json:"code"`    // The HTTP status code.
	Status  string `json:"status"`  // The HTTP response status (error or success).
	Message string `json:"message"` // A short description of the error.
}

// decodeError decodes an error from an io.Reader.
func (c *Client) decodeError(resp *http.Response) error {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(respBody) == 0 {
		return fmt.Errorf("HTTP %d: %s (body empty)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	buf := bytes.NewBuffer(respBody)
	var e Error
	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("couldn't decode error: [%s]", respBody)
	}

	return fmt.Errorf("error %v: %s ", e.Code, e.Message)
}

}
