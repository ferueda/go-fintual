package fintual

import (
	"context"
)

const (
	banksEndpoint = "/banks"
)

// BanksService handles communication with the Banks related
// methods of the Fintual API.
//
// Fintual API docs: https://fintual.cl/api-docs
type BanksService service

type Bank struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes BankAttributes `json:"attributes"`
}

type BankAttributes struct {
	Name string `json:"name"`
}

// BankListParams specifies the optional parameters to the
// BanksService.ListAll method.
type BankListParams struct {
	Query string `url:"q,omitempty"` // For filtering results
}

// ListAll lists all Banks. Receives a params argument
// with a Query property for filtering Banks by the Name attribute.
//
// Endpoint: GET /banks
func (s *BanksService) ListAll(ctx context.Context, params *BankListParams) ([]*Bank, error) {
	url := s.client.baseURL.String() + banksEndpoint
	url, err := addParams(url, params)
	if err != nil {
		return nil, err
	}

	var banks struct {
		Data []*Bank `json:"data"`
	}

	err = s.client.get(ctx, url, &banks)
	if err != nil {
		return nil, err
	}

	return banks.Data, nil
}
