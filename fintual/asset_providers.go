package fintual

import "context"

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

// GetAssetProviders retrieves Fintual's asset providers
func (c *Client) GetAssetProviders(ctx context.Context) ([]*AssetProvider, error) {
	var ap struct {
		Data []*AssetProvider `json:"data"`
	}

	err := c.get(ctx, baseURL+assetProvidersEndpoint, &ap)
	if err != nil {
		return nil, err
	}

	return ap.Data, nil
}
