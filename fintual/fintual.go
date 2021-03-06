// Package fintual provides utilties for interfacing
// with the Fintual API.
package fintual

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
)

// Version is the version of this library.
const Version = "1.0.0"

const (
	baseURL = "https://fintual.cl/api"
)

type service struct {
	client *Client
}

type Client struct {
	http        *http.Client // HTTP client used to communicate with the API.
	baseURL     *url.URL     // Base URL for API requests
	userEmail   string       // User's email used for methods which require authentication
	accessToken string       // Access token used for methods which require authentication

	// Services used for talking to different parts of the Fintual API.
	AssetProviders   *AssetProvidersService
	Banks            *BanksService
	ConceptualAssets *ConceptualAssetsService
	Goals            *GoalsService
	RealAssets       *RealAssetsService
}

// NewClient returns a new Fintual API client.
// If a nil httpClient is provided, a new http.Client will be used.
//
// To use API methods which require authentication, you must call
// the Client.Authenticate method with valid credentials.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: time.Minute}
	}
	baseURL, _ := url.Parse(baseURL)

	c := &Client{http: httpClient, baseURL: baseURL}

	c.AssetProviders = &AssetProvidersService{client: c}
	c.Banks = &BanksService{client: c}
	c.ConceptualAssets = &ConceptualAssetsService{client: c}
	c.Goals = &GoalsService{client: c}
	c.RealAssets = &RealAssetsService{client: c}
	return c
}

// setAccessToken sets the given token to the current Fintual client.
func (c *Client) setAccessToken(token string) {
	c.accessToken = token
}

// setUserEmail sets the given email to the current Fintual client.
func (c *Client) setUserEmail(token string) {
	c.userEmail = token
}

// addParams adds the parameters in params as URL query parameters to s. params
// must be a struct whose fields may contain "url" tags.
func addParams(s string, params interface{}) (string, error) {
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(params)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
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

// newRequest creates a new API request with context. If specified,
// the value pointed to by body is JSON encoded and included in the request body.
func (c *Client) newRequest(ctx context.Context, method, url string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// send makes a request to the API, the response body will be
// unmarshalled into v.
func (c *Client) send(req *http.Request, v interface{}) error {
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return c.decodeError(resp)
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

// get makes a GET request to the given url. The response body will be
// unmarshalled into v.
func (c *Client) get(ctx context.Context, url string, v interface{}) error {
	req, err := c.newRequest(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	err = c.send(req, v)
	if err != nil {
		return err
	}

	return nil
}

// post makes a POST request to the given url. The response body will be
// unmarshalled into v.
func (c *Client) post(ctx context.Context, url string, body, v interface{}) error {
	req, err := c.newRequest(ctx, "POST", url, body)
	if err != nil {
		return err
	}

	err = c.send(req, v)
	if err != nil {
		return err
	}

	return nil
}

// getWithAuth makes a GET request with authentication credentials
// to the given url. The response body will be unmarshalled into v.
func (c *Client) getWithAuth(ctx context.Context, url string, v interface{}) error {
	if c.accessToken == "" || c.userEmail == "" {
		return errors.New("client not authenticated, call Client.Authenticate with valid credentials")
	}

	req, err := c.newRequest(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("user_email", c.userEmail)
	q.Add("user_token", c.accessToken)
	req.URL.RawQuery = q.Encode()

	err = c.send(req, v)
	if err != nil {
		return err
	}

	return nil
}
