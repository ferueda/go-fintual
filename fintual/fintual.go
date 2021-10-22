// Package fintual provides utilties for interfacing
// with the Fintual API.
package fintual

import (
	"net/http"
	"net/url"
)

// Version is the version of this library.
const Version = "1.0.0"

const (
	baseURL = "https://fintual.cl/api"
)

type Client struct {
	client      *http.Client // HTTP client used to communicate with the API.
	baseURL     *url.URL     // Base URL for API requests
	accessToken string       // Access token used for methods which require authentication
}

// NewClient returns a new Fintual API client.
// To use API methods which require authentication, you must first call
// the provided Authenticate method with valid credentials.
func NewClient() *Client {
	baseURL, _ := url.Parse(baseURL)
	return &Client{client: &http.Client{}, baseURL: baseURL}
}
