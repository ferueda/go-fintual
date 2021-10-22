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

type accessToken struct {
	Type       string         `json:"type"`
	Attributes authAttributes `json:"attributes"`
}

type authAttributes struct {
	Token string `json:"token"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Authenticate tries to retrieve a user access token from the
// Fintual access_tokens endpoint and sets it to the current Fintual client
func (c *Client) Authenticate(email, password string) error {
	reqData := struct {
		User Credentials `json:"user"`
	}{User: Credentials{Email: email, Password: password}}

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return err
	}
	resp, err := c.http.Post(baseURL+accessTokenEndpoint, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return c.decodeError(resp)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var data struct {
		Data accessToken `json:"data"`
	}
	json.Unmarshal(respBody, &data)

	if data.Data.Type != "access_token" || len(data.Data.Attributes.Token) == 0 {
		return errors.New("fitual auth failed - didn't get access token")
	}

	c.accessToken = data.Data.Attributes.Token
	return nil
}
