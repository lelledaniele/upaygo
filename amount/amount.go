package amount

import "github.com/lelledaniele/upaygo/currency"

type Amount interface {
	GetAmount() int
	GetCurrency() currency.Currency
}

type a struct {
	A int               `json:"amount"`
	C currency.Currency `json:"currency"`
}

// GetAmount exposes a.A value
func (a *a) GetAmount() int {
	return a.A
}

// GetCurrency exposes a.C value
func (a *a) GetCurrency() currency.Currency {
	return a.C
}

// New returns a new instance of a
func New(v int, cs string) (Amount, error) {
	c, e := currency.New(cs)
	if e != nil {
		return nil, e
	}

	return &a{v, c}, nil
}
