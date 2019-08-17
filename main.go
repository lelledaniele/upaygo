package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	appconfig "github.com/lelledaniele/upaygo/config"
	apprestintentcreate "github.com/lelledaniele/upaygo/controller/rest/intent/create"

	_ "github.com/lelledaniele/upaygo/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

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
const configFile = "config.json"

// @title uPayment in GO
// @version 1.0.0
// @description Microservice to manage payment
// @license.name MIT
func main() {
	fc, e := os.Open(configFile)
	if e != nil {
		fmt.Printf("Impossible to get configuration file: %v\n", e)
		os.Exit(1)
	}
	defer fc.Close()

	e = appconfig.ImportConfig(fc)
	if e != nil {
		fmt.Printf("Error durring file config import: %v", e)
		os.Exit(1)
	}

	s := appconfig.GetServerConfig()

	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(s.GetURI()+"/swagger/doc.json"),
	))
	http.HandleFunc(apprestintentcreate.URL, apprestintentcreate.Handler)

	log.Fatal(http.ListenAndServe(":"+s.GetPort(), nil))
}
