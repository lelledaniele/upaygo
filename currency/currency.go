package appcurrency

import (
	"fmt"
	"strings"
)

type Currency interface {
	GetISO4217() string
}

type c struct {
	ISO4217 string `json:"ISO_4217"` // 3 Char symbol - ISO 4217
}

// GetISO4217 exposes c.ISO4217 value
func (c *c) GetISO4217() string {
	return c.ISO4217
}

// New returns a new instance of c
func New(ISO4217 string) (Currency, error) {
	if len(ISO4217) != 3 {
		return nil, fmt.Errorf("'%v' is not a currency ISO 4217 format", ISO4217)
	}

	return &c{ISO4217: strings.ToUpper(ISO4217)}, nil
}
