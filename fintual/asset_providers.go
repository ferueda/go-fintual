package fintual

import (
	"context"
	"fmt"
)

const (
	assetProvidersEndpoint = "/asset_providers"
)

// AssetProvidersService handles communication with the
// Asset Providers related methods of the Fintual API.
//
// Fintual API docs: https://fintual.cl/api-docs
type AssetProvidersService service

type AssetProvider struct {
	ID         string                  `json:"id"`
	Type       string                  `json:"type"`
	Attributes AssetProviderAttributes `json:"attributes"`
}

type AssetProviderAttributes struct {
	Name string `json:"name"`
}

// ListAll lists all asset providers.
//
// Endpoint: GET /asset_providers
func (s *AssetProvidersService) ListAll(ctx context.Context) ([]*AssetProvider, error) {
	url := s.client.baseURL.String() + assetProvidersEndpoint
	var ap struct {
		Data []*AssetProvider `json:"data"`
	}

	err := s.client.get(ctx, url, &ap)
	if err != nil {
		return nil, err
	}

	return ap.Data, nil
}

// Get retrieves a single asset provider.
//
// Endpoint: GET /asset_providers/:id
func (s *AssetProvidersService) Get(ctx context.Context, id string) (*AssetProvider, error) {
	url := fmt.Sprintf("%s/%s", s.client.baseURL.String()+assetProvidersEndpoint, id)
	var ap struct {
		Data *AssetProvider `json:"data"`
	}

	err := s.client.get(ctx, url, &ap)
	if err != nil {
		return nil, err
	}

	return ap.Data, nil
}
