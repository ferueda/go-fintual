package fintual

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

const (
	realAssetsEndpoint   = "/real_assets"
	expenseRatioEndpoint = "/expense_ratio"
	daysEndpoint         = "/days"
)

// RealAssetsService handles communication with the
// Real Assets related methods of the Fintual API.
//
// Fintual API docs: https://fintual.cl/api-docs
type RealAssetsService service

type RealAsset struct {
	ID         string              `json:"id"`
	Type       string              `json:"type"`
	Attributes RealAssetAttributes `json:"attributes"`
}

type RealAssetAttributes struct {
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	Serie             string      `json:"serie"`
	StartDate         string      `json:"start_date"`
	EndDate           interface{} `json:"end_date"`
	PreviousAssetID   interface{} `json:"previous_asset_id"`
	LastDay           LastDay     `json:"last_day"`
	ConceptualAssetID int         `json:"conceptual_asset_id"`
}

type LastDay struct {
	FixedManagementFee     float64 `json:"fixed_management_fee"`
	IvaExclusiveExpenses   float64 `json:"iva_exclusive_expenses"`
	IvaInclusiveExpenses   float64 `json:"iva_inclusive_expenses"`
	NetAssetValue          float64 `json:"net_asset_value"`
	PurchaseFee            float64 `json:"purchase_fee"`
	RedemptionFee          float64 `json:"redemption_fee"`
	TotalAssets            float64 `json:"total_assets"`
	TotalNetAssets         float64 `json:"total_net_assets"`
	VariableManagementFee  float64 `json:"variable_management_fee"`
	FixedFee               float64 `json:"fixed_fee"`
	NewShares              float64 `json:"new_shares"`
	OutstandingShares      float64 `json:"outstanding_shares"`
	RedeemedShares         float64 `json:"redeemed_shares"`
	InstitutionalInvestors float64 `json:"institutional_investors"`
	Shareholders           float64 `json:"shareholders"`
	Date                   string  `json:"date"`
}

// Get retrieves a single Real Asset.
//
// Endpoint: GET /real_assets/:id
func (s *RealAssetsService) Get(ctx context.Context, id string) (*RealAsset, error) {
	url := fmt.Sprintf("%s/%s", s.client.baseURL.String()+realAssetsEndpoint, id)

	var ra struct {
		Data *RealAsset `json:"data"`
	}

	err := s.client.get(ctx, url, &ra)
	if err != nil {
		return nil, err
	}

	return ra.Data, nil
}

type ExpenseRationRealAsset struct {
	ID         string                           `json:"id"`
	Type       string                           `json:"type"`
	Attributes ExpenseRationRealAssetAttributes `json:"attributes"`
}

type ExpenseRationRealAssetAttributes struct {
	ExpenseRatio float64 `json:"expense_ratio"`
}

// GetExpenseRatio retrieves the Expense Ratio for a single Real Asset.
//
// Endpoint: GET /real_assets/:id/expense_ratio
func (s *RealAssetsService) GetExpenseRatio(ctx context.Context, id string) (*ExpenseRationRealAsset, error) {
	url := fmt.Sprintf("%s/%s%s", s.client.baseURL.String()+realAssetsEndpoint, id, expenseRatioEndpoint)

	var ra struct {
		Data *ExpenseRationRealAsset `json:"data"`
	}

	err := s.client.get(ctx, url, &ra)
	if err != nil {
		return nil, err
	}

	return ra.Data, nil
}

type RealAssetDay struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes RealAssetDayAttributes `json:"attributes"`
}

type RealAssetDayAttributes struct {
	Date                       string  `json:"date"`
	Price                      float64 `json:"price"`
	FixedManagementFee         float64 `json:"fixed_management_fee"`
	FixedManagementFeeType     string  `json:"fixed_management_fee_type"`
	IvaExclusiveExpenses       float64 `json:"iva_exclusive_expenses"`
	IvaExclusiveExpensesType   string  `json:"iva_exclusive_expenses_type"`
	IvaInclusiveExpenses       float64 `json:"iva_inclusive_expenses"`
	IvaInclusiveExpensesType   string  `json:"iva_inclusive_expenses_type"`
	NetAssetValue              float64 `json:"net_asset_value"`
	NetAssetValueType          string  `json:"net_asset_value_type"`
	PurchaseFee                float64 `json:"purchase_fee"`
	PurchaseFeeType            string  `json:"purchase_fee_type"`
	RedemptionFee              float64 `json:"redemption_fee"`
	RedemptionFeeType          string  `json:"redemption_fee_type"`
	TotalAssets                float64 `json:"total_assets"`
	TotalAssetsType            string  `json:"total_assets_type"`
	TotalNetAssets             float64 `json:"total_net_assets"`
	TotalNetAssetsType         string  `json:"total_net_assets_type"`
	VariableManagementFee      float64 `json:"variable_management_fee"`
	VariableManagementFeeType  string  `json:"variable_management_fee_type"`
	FixedFee                   float64 `json:"fixed_fee"`
	FixedFeeType               string  `json:"fixed_fee_type"`
	NewShares                  float64 `json:"new_shares"`
	NewSharesType              string  `json:"new_shares_type"`
	OutstandingShares          float64 `json:"outstanding_shares"`
	OutstandingSharesType      string  `json:"outstanding_shares_type"`
	RedeemedShares             float64 `json:"redeemed_shares"`
	RedeemedSharesType         string  `json:"redeemed_shares_type"`
	InstitutionalInvestors     float64 `json:"institutional_investors"`
	InstitutionalInvestorsType string  `json:"institutional_investors_type"`
	Shareholders               float64 `json:"shareholders"`
	ShareholdersType           string  `json:"shareholders_type"`
}

// GetDay retrieves a Real Asset Day. Receives a Real Asset ID
// and a string date with format YYYY-MM-DD.
//
// Endpoint: GET /real_assets/:id/days
func (s *RealAssetsService) GetDay(ctx context.Context, id string, date string) ([]*RealAssetDay, error) {
	if date == "" || len(strings.Split(date, "-")) != 3 {
		return nil, errors.New("received malformatted or zero value date")
	}

	url := fmt.Sprintf("%s/%s%s?date=%s", s.client.baseURL.String()+realAssetsEndpoint, id, daysEndpoint, date)

	var rad struct {
		Data []*RealAssetDay `json:"data"`
	}

	err := s.client.get(ctx, url, &rad)
	if err != nil {
		return nil, err
	}

	return rad.Data, nil
}

// ListDaysByDates lists Real Asset Days. Receives a Real Asset ID
// and a from and to string dates with format YYYY-MM-DD.
//
// Endpoint: GET /real_assets/:id/days
func (s *RealAssetsService) ListDaysByDates(ctx context.Context, id, from, to string) ([]*RealAssetDay, error) {
	if from == "" || to == "" || len(strings.Split(from, "-")) != 3 || len(strings.Split(to, "-")) != 3 {
		return nil, errors.New("received malformatted or zero value dates")
	}

	url := fmt.Sprintf("%s/%s%s?from_date=%s&to_date=%s", s.client.baseURL.String()+realAssetsEndpoint, id, daysEndpoint, from, to)

	var rad struct {
		Data []*RealAssetDay `json:"data"`
	}

	err := s.client.get(ctx, url, &rad)
	if err != nil {
		return nil, err
	}

	return rad.Data, nil
}

type ConceptualAssetRealAsset struct {
	ID         string                             `json:"id"`
	Type       string                             `json:"type"`
	Attributes ConceptualAssetRealAssetAttributes `json:"attributes"`
}

type ConceptualAssetRealAssetAttributes struct {
	Name              string                          `json:"name"`
	Symbol            string                          `json:"symbol"`
	Serie             string                          `json:"serie"`
	StartDate         string                          `json:"start_date"`
	EndDate           string                          `json:"end_date"`
	PreviousAssetID   string                          `json:"previous_asset_id"`
	LastDay           ConceptualAssetRealAssetLastDay `json:"last_day"`
	ConceptualAssetID int                             `json:"conceptual_asset_id"`
}

type ConceptualAssetRealAssetLastDay struct {
	Rate float64 `json:"rate"`
	Date string  `json:"date"`
}

// ListByConceptualAsset lists all Real Assets
// of a given Conceptual Asset.
//
// Endpoint: GET /conceptual_assets/:id/real_assets
func (s *RealAssetsService) ListByConceptualAsset(ctx context.Context, id string) ([]*ConceptualAssetRealAsset, error) {
	url := fmt.Sprintf("%s/%s%s", s.client.baseURL.String()+conceptualAssetsEndpoint, id, realAssetsEndpoint)

	var d struct {
		Data []*ConceptualAssetRealAsset `json:"data"`
	}

	err := s.client.get(ctx, url, &d)
	if err != nil {
		return nil, err
	}

	return d.Data, nil
}
