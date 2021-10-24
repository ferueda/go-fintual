package fintual

import (
	"context"
	"fmt"
)

const (
	assetProvidersEndpoint = "/asset_providers"
)

type AssetProvider struct {
	ID         string                  `json:"id"`
	Type       string                  `json:"type"`
	Attributes AssetProviderAttributes `json:"attributes"`
}

type AssetProviderAttributes struct {
	Name string `json:"name"`
}

// GetAssetProvider retrieves a single asset provider.
//
// Endpoint: GET /asset_providers/:id
func (c *Client) GetAssetProvider(ctx context.Context, id string) (*AssetProvider, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL.String()+assetProvidersEndpoint, id)
	var ap struct {
		Data *AssetProvider `json:"data"`
	}

	err := c.get(ctx, url, &ap)
	if err != nil {
		return nil, err
	}

	return ap.Data, nil
}

// ListAssetProviders lists all asset providers.
//
// Endpoint: GET /asset_providers
func (c *Client) ListAssetProviders(ctx context.Context) ([]*AssetProvider, error) {
	url := c.baseURL.String() + assetProvidersEndpoint
	var ap struct {
		Data []*AssetProvider `json:"data"`
	}

	err := c.get(ctx, url, &ap)
	if err != nil {
		return nil, err
	}

	return ap.Data, nil
}
