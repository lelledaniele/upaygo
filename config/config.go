package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configFile = "config.json"

//{
// "stripe": {
//    "api_keys": {
//	  "EUR": {
//		"pk_key": "pk_EUR",
//		"sk_key": "sk_EUR"
//	  },
//	  "default": {
//        "pk_key": "pk_DEFAULT",
//        "sk_key": "sk_DEFAULT"
//      }
//    }
//  },
//  "server": {
//    "port": "8080"
//  }
//}
var s config

// Reads the config json file and save the content in s global var
func init() {
	fc, e := ioutil.ReadFile(configFile)
	if e != nil {
		fmt.Printf("Impossible to get configuration file: %v", e)
		os.Exit(1)
	}

	e = json.Unmarshal(fc, &s)
	if e != nil {
		fmt.Printf("Impossible to unmarshal configuration file: %v", e)
		os.Exit(1)
	}
}

// public properties needed for json.Unmarshal
type config struct {
	Stripe apiKeys `json:"stripe"`
	Server server  `json:"server"`
}
