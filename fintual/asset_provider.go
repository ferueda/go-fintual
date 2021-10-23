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

// GetAssetProvider retrieves a single Fintual asset provider
func (c *Client) GetAssetProvider(ctx context.Context, id string) (*AssetProvider, error) {
	var ap struct {
		Data *AssetProvider `json:"data"`
	}

	url := fmt.Sprintf("%s/%s", c.baseURL.String()+assetProvidersEndpoint, id)
	err := c.get(ctx, url, &ap)
	if err != nil {
		return nil, err
	}

	return ap.Data, nil
}

// GetAssetProviders retrieves Fintual's asset providers list
func (c *Client) GetAssetProviders(ctx context.Context) ([]*AssetProvider, error) {
	var ap struct {
		Data []*AssetProvider `json:"data"`
	}

	url := c.baseURL.String() + assetProvidersEndpoint
	err := c.get(ctx, url, &ap)
	if err != nil {
		return nil, err
	}

	return ap.Data, nil
}
