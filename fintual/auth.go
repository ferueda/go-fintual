package fintual

import (
	"context"
	"errors"
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
func (c *Client) Authenticate(ctx context.Context, email, password string) error {
	reqBody := struct {
		User Credentials `json:"user"`
	}{User: Credentials{Email: email, Password: password}}

	var data struct {
		Data accessToken `json:"data"`
	}
	url := c.baseURL.String() + accessTokenEndpoint
	err := c.post(ctx, url, reqBody, &data)
	if err != nil {
		return err
	}

	if data.Data.Type != "access_token" || len(data.Data.Attributes.Token) == 0 {
		return errors.New("fitual auth failed - didn't get access token")
	}

	c.setAccessToken(data.Data.Attributes.Token)
	return nil
}
