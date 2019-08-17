package appconfig

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

var s config

// ImportConfig reads the config json reader and save the content in s global var
func ImportConfig(r io.Reader) error {
	b, e := ioutil.ReadAll(r)
	if e != nil {
		return fmt.Errorf("Impossible to read configuration: %v\n", e)
	}

	s = config{} // Reset existing configs
	e = json.Unmarshal(b, &s)
	if e != nil {
		return fmt.Errorf("Impossible to unmarshal configuration: %v\n", e)
	}

	return nil
}

// public properties needed for json.Unmarshal
type config struct {
	Stripe apiKeys `json:"stripe"`
	Server server  `json:"server"`
}
