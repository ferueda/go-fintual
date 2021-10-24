package fintual

import (
	"context"
	"fmt"
)

const (
	conceptualAssetsEndpoint = "/conceptual_assets"
)

// ConceptualAssetsService handles communication with the
// Conceptual Assets related methods of the Fintual API.
//
// Fintual API docs: https://fintual.cl/api-docs
type ConceptualAssetsService service

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

// ListAll lists all conceptual assets. Receives a params argument
// with Name and/or Run properties for filtering Conceptual Assets
// by the Name and Run attributes.
//
// Endpoint: GET /conceptual_assets
func (s *ConceptualAssetsService) ListAll(ctx context.Context, params *ConceptualAssetListParams) ([]*ConceptualAsset, error) {
	url := s.client.baseURL.String() + conceptualAssetsEndpoint
	url, err := addParams(url, params)
	if err != nil {
		return nil, err
	}

	var ca struct {
		Data []*ConceptualAsset `json:"data"`
	}

	err = s.client.get(ctx, url, &ca)
	if err != nil {
		return nil, err
	}

	return ca.Data, nil
}

// Get retrieves a single conceptual asset.
//
// Endpoint: GET /conceptual_assets/:id
func (s *ConceptualAssetsService) Get(ctx context.Context, id string) (*ConceptualAsset, error) {
	url := fmt.Sprintf("%s/%s", s.client.baseURL.String()+conceptualAssetsEndpoint, id)

	var ca struct {
		Data *ConceptualAsset `json:"data"`
	}

	err := s.client.get(ctx, url, &ca)
	if err != nil {
		return nil, err
	}

	return ca.Data, nil
}

// ListByAssetProvider lists all conceptual assets
// of a given Asset Provider. Receives a params argument
// with Name and/or Run properties for filtering Conceptual Assets
// by the Name and Run attributes.
//
// Endpoint: GET /asset_providers/:id/conceptual_assets
func (s *ConceptualAssetsService) ListByAssetProvider(ctx context.Context, id string, params *ConceptualAssetListParams) ([]*ConceptualAsset, error) {
	url := fmt.Sprintf("%s/%s/%s", s.client.baseURL.String()+assetProvidersEndpoint, id, conceptualAssetsEndpoint)
	url, err := addParams(url, params)
	if err != nil {
		return nil, err
	}

	var ca struct {
		Data []*ConceptualAsset `json:"data"`
	}

	err = s.client.get(ctx, url, &ca)
	if err != nil {
		return nil, err
	}

	return ca.Data, nil
}
