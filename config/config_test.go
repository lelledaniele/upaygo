package config

import (
	"encoding/json"
	"testing"
)

const (
	confWithDefault = `
{
  "stripe": {
    "api_keys": {
	  "EUR": {
		"pk_key": "pk_EUR",
		"sk_key": "sk_EUR"
	  },
	  "default": {
        "pk_key": "pk_DEFAULT",
        "sk_key": "sk_DEFAULT"
      }
    }
  }
}
`
	confWithoutDefault = `
{
  "stripe": {
    "api_keys": {
      "EUR": {
        "pk_key": "pk_EUR",
        "sk_key": "sk_EUR"
      }
    }
  }
}
`
)

func TestDefaultStripeAPIConfig(t *testing.T) {
	s = config{}
	_ = json.Unmarshal([]byte(confWithDefault), &s)
	got, e := GetStripeAPIConfigByCurrency("NOT_FOUND_CURRENCY")

	if e != nil {
		t.Errorf("error during the retrieve of Stripe API config by inexistent currency: %v", e)
	}

	if got.GetPK() != "pk_DEFAULT" || got.GetSK() != "sk_DEFAULT" {
		t.Errorf("Stripe API config for an inexistent currency does not return the default value")
	}
}

func TestWithoutDefaultStripeAPIConfig(t *testing.T) {
	s = config{}
	_ = json.Unmarshal([]byte(confWithoutDefault), &s)
	_, e := GetStripeAPIConfigByCurrency("NOT_FOUND_CURRENCY")

	if e == nil {
		t.Error("configuration without Stripe default API keys must return an error")
	}
}

func TestCurrencyStripeAPIConfig(t *testing.T) {
	s = config{}
	_ = json.Unmarshal([]byte(confWithDefault), &s)
	got, e := GetStripeAPIConfigByCurrency("EUR")

	if e != nil {
		t.Errorf("error during the retrieve of Stripe API config by EUR currency: %v", e)
	}

	if got.GetPK() != "pk_EUR" || got.GetSK() != "sk_EUR" {
		t.Errorf("incorrect Stripe API config for EUR currency, got %v and %v, want %v %v", got.GetPK(), got.GetSK(), "pk_EUR", "sk_EUR")
	}
}

func TestLowercaseCurrencyStripeAPIConfig(t *testing.T) {
	s = config{}
	_ = json.Unmarshal([]byte(confWithDefault), &s)
	got, e := GetStripeAPIConfigByCurrency("eur")

	if e != nil {
		t.Errorf("error during the retrieve of Stripe API config by eur currency: %v", e)
	}

	if got.GetPK() != "pk_EUR" || got.GetSK() != "sk_EUR" {
		t.Errorf("incorrect Stripe API config for eur currency, got %v and %v, want %v %v", got.GetPK(), got.GetSK(), "pk_EUR", "sk_EUR")
	}
}
