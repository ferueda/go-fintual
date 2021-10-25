package fintual

import (
	"context"
	"fmt"
)

const (
	goalsEndpoint = "/goals"
)

// GoalsService handles communication with the
// Goals related methods of the Fintual API.
//
// Fintual API docs: https://fintual.cl/api-docs
type GoalsService service

type Goal struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes GoalAttributes `json:"attributes"`
}

type GoalAttributes struct {
	Name                   string       `json:"name"`
	NameWithoutSuffix      string       `json:"name_without_suffix"`
	NetAssetValue          float64      `json:"nav"`
	CreatedAt              string       `json:"created_at"`
	Timeframe              int          `json:"timeframe"`
	Deposited              float64      `json:"deposited"`
	Hidden                 bool         `json:"hidden"`
	Profit                 float64      `json:"profit"`
	Investments            []Investment `json:"investments"`
	PublicLink             interface{}  `json:"public_link"`
	ParamID                int64        `json:"param_id"`
	GoalType               string       `json:"goal_type"`
	TranslatedGoalType     string       `json:"translated_goal_type"`
	Regime                 interface{}  `json:"regime"`
	Completed              bool         `json:"completed"`
	HasAnyWithdrawals      bool         `json:"has_any_withdrawals"`
	EligibleForDeposits    bool         `json:"eligible_for_deposits"`
	EligibleForInternalMlt bool         `json:"eligible_for_internal_mlt"`
	MonthlyDeposit         float64      `json:"monthly_deposit"`
	SimulatedDeposit       float64      `json:"simulated_deposit"`
	FundsSource            interface{}  `json:"funds_source"`
	FundsSourceDescription interface{}  `json:"funds_source_description"`
	NotNetDeposited        float64      `json:"not_net_deposited"`
	Withdrawn              float64      `json:"withdrawn"`
	GroupGoalID            interface{}  `json:"group_goal_id"`
}

type Investment struct {
	Weight  float64 `json:"weight"`
	AssetID int     `json:"asset_id"`
}

// ListAll lists all Goals for the authenticated user.
// Requires authentication by calling Client.Authenticate first.
//
// Endpoint: GET /goals
func (s *GoalsService) ListAll(ctx context.Context) ([]*Goal, error) {
	url := s.client.baseURL.String() + goalsEndpoint
	var g struct {
		Data []*Goal `json:"data"`
	}

	err := s.client.getWithAuth(ctx, url, &g)
	if err != nil {
		return nil, err
	}

	return g.Data, nil
}

// Get retrieves a specific goal.
// Requires authentication by calling Client.Authenticate first.
//
// Endpoint: GET /goals/:id
func (s *GoalsService) Get(ctx context.Context, id string) (*Goal, error) {
	url := fmt.Sprintf("%s/%s", s.client.baseURL.String()+goalsEndpoint, id)
	var g struct {
		Data *Goal `json:"data"`
	}

	err := s.client.getWithAuth(ctx, url, &g)
	if err != nil {
		return nil, err
	}

	return g.Data, nil
}
