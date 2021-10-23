package fintual

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
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
	reqData := struct {
		User Credentials `json:"user"`
	}{User: Credentials{Email: email, Password: password}}

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	url := c.baseURL.String() + accessTokenEndpoint
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(req)
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
