package appconfig

import (
	"errors"
	"strings"
)

// GetStripeAPIConfigByCurrency returns the API keys by c currency
// returns default API keys if c not in b
func GetStripeAPIConfigByCurrency(c string) (APIConfig, error) {
	r, f := s.Stripe.Keys[strings.ToUpper(c)]
	if f {
		return &r, nil
	}

	r, f = s.Stripe.Keys["default"]
	if f {
		return &r, nil
	}

	return nil, errors.New("default API keys not found")
}

// APIConfig exposes the PK and SK, without exposing the struct (currencyAPIConfig)
type APIConfig interface {
	GetPK() string
	GetSK() string
}

// public properties needed for json.Unmarshal
type apiKeys struct {
	Keys map[string]currencyAPIConfig `json:"api_keys"`
}

// public properties needed for json.Unmarshal
type currencyAPIConfig struct {
	PK string `json:"pk_key"`
	SK string `json:"sk_key"`
}

// GetSK exposes currencyAPIConfig.PK
func (s *currencyAPIConfig) GetPK() string {
	return s.PK
}

// GetSK exposes currencyAPIConfig.SK
func (s *currencyAPIConfig) GetSK() string {
	return s.SK
}
