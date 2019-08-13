package customer

type Customer interface {
	GetGatewayReference() string
}

type c struct {
	R string `json:"gateway_reference"` // Gateway reference
}

// GetGatewayReference exposes c.R value
func (c *c) GetGatewayReference() string {
	return c.R
}

// New returns a new instance of c
func New(r string) Customer {
	return &c{r}
}
