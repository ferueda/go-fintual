package fintual

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

const (
	accessTokenEndpoint = "/access_tokens"
)

type authResponse struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			Token string `json:"token"`
		} `json:"attributes"`
	} `json:"data"`
}

// Authenticate tries to retrieve a user access token from the
// Fintual access_tokens endpoint and sets it to the current Fintual client
func (c *Client) Authenticate(email, password string) error {
	creds := map[string]map[string]string{"user": {"email": email, "password": password}}
	body, err := json.Marshal(creds)
	if err != nil {
		return err
	}
	resp, err := c.client.Post(baseURL+accessTokenEndpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		switch resp.StatusCode {
		case 401:
			return errors.New("fitual auth failed - invalid credentials")
		default:
			return errors.New("fitual auth failed - unknown error")
		}
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var respObj authResponse
	json.Unmarshal(data, &respObj)

	if respObj.Data.Type != "access_token" || len(respObj.Data.Attributes.Token) == 0 {
		return errors.New("fitual auth failed - didn't get access token")
	}

	c.accessToken = respObj.Data.Attributes.Token
	return nil
}
