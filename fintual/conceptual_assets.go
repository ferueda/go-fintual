package fintual

import (
	"context"
	"fmt"
)

const (
	conceptualAssetsEndpoint = "/conceptual_assets"
)

type ConceptualAsset struct {
	ID         string                    `json:"id"`
	Type       string                    `json:"type"`
	Attributes ConceptualAssetAttributes `json:"attributes"`
}

type ConceptualAssetAttributes struct {
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	Category   string `json:"category"`
	Currency   string `json:"currency"`
	MaxScale   int    `json:"max_scale"`
	Run        string `json:"run"`
	DataSource string `json:"data_source"`
}

// ConceptualAssetListParams specifies the optional parameters to the
// ListConceptualAssets method.
type ConceptualAssetListParams struct {
	Name string `url:"name,omitempty"` // For filtering results by name
	Run  string `url:"run,omitempty"`  // For filtering results by run identifier
}

// GetConceptualAsset retrieves a single conceptual asset.
//
// Endpoint: GET /conceptual_assets/:id
func (c *Client) GetConceptualAsset(ctx context.Context, id string) (*ConceptualAsset, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL.String()+conceptualAssetsEndpoint, id)
	var ca struct {
		Data *ConceptualAsset `json:"data"`
	}

	err := c.get(ctx, url, &ca)
	if err != nil {
		return nil, err
	}

	return ca.Data, nil
}

// ListConceptualAssets lists all conceptual assets. Receives a params argument
// with Name and/or Run properties for filtering Conceptual Assets
// by the Name and Run attributes.
//
// Endpoint: GET /conceptual_assets
func (c *Client) ListConceptualAssets(ctx context.Context, params *ConceptualAssetListParams) ([]*ConceptualAsset, error) {
	url := c.baseURL.String() + conceptualAssetsEndpoint
	url, err := addParams(url, params)
	if err != nil {
		return nil, err
	}

	var ca struct {
		Data []*ConceptualAsset `json:"data"`
	}

	err = c.get(ctx, url, &ca)
	if err != nil {
		return nil, err
	}

	return ca.Data, nil
}

// ListAssetProviderConceptualAssets lists all conceptual assets
// of a given Asset Provider.
//
// Endpoint: GET /asset_providers/:id/conceptual_assets
func (c *Client) ListAssetProviderConceptualAssets(ctx context.Context, id string) ([]*ConceptualAsset, error) {
	url := fmt.Sprintf("%s/%s/%s", c.baseURL.String()+assetProvidersEndpoint, id, conceptualAssetsEndpoint)
	var ca struct {
		Data []*ConceptualAsset `json:"data"`
	}

	err := c.get(ctx, url, &ca)
	if err != nil {
		return nil, err
	}

	return ca.Data, nil
}
