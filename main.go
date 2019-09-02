package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	appconfig "github.com/lelledaniele/upaygo/config"
	apprestintentcancel "github.com/lelledaniele/upaygo/controller/rest/intent/cancel"
	apprestintentcapture "github.com/lelledaniele/upaygo/controller/rest/intent/capture"
	apprestintentconfirm "github.com/lelledaniele/upaygo/controller/rest/intent/confirm"
	apprestintentcreate "github.com/lelledaniele/upaygo/controller/rest/intent/create"
	apprestintentget "github.com/lelledaniele/upaygo/controller/rest/intent/get"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
	_ "github.com/lelledaniele/upaygo/docs"
)

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

	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	r.HandleFunc(apprestintentget.URL, apprestintentget.Handler)
	r.HandleFunc(apprestintentcreate.URL, apprestintentcreate.Handler)
	r.HandleFunc(apprestintentconfirm.URL, apprestintentconfirm.Handler)
	r.HandleFunc(apprestintentcapture.URL, apprestintentcapture.Handler)
	r.HandleFunc(apprestintentcancel.URL, apprestintentcancel.Handler)

	log.Fatal(http.ListenAndServe(":"+s.GetPort(), r))
}
