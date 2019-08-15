package apppaymentsource

type Source interface {
	GetGatewayReference() string
}

type ps struct {
	R string `json:"gateway_reference"` // Gateway reference
}

// GetGatewayReference exposes ps.R value
func (ps *ps) GetGatewayReference() string {
	return ps.R
}

// New returns a new instance of ps
func New(r string) Source {
	return &ps{r}
}
