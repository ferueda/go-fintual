package fintual

import (
	"context"
)

const (
	banksEndpoint = "/banks"
)

type Bank struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes BankAttributes `json:"attributes"`
}

type BankAttributes struct {
	Name string `json:"name"`
}

// BankListParams specifies the optional parameters to the
// ListBanks method.
type BankListParams struct {
	ListParams
}

// ListBanks retrieves a list of all Banks. Receives an optional query param Query
// which filter banks based on the bank's Name attribute.
//
// Endpoint: GET /banks
func (c *Client) ListBanks(ctx context.Context, params *BankListParams) ([]Bank, error) {
	url := c.baseURL.String() + banksEndpoint
	url, err := addParams(url, params)
	if err != nil {
		return nil, err
	}

	var banks struct {
		Data []Bank `json:"data"`
	}

	err = c.get(ctx, url, &banks)
	if err != nil {
		return nil, err
	}

	return banks.Data, nil
}
