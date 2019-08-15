package main

import (
	"log"
	"net/http"

	restintentcreate "github.com/lelledaniele/upaygo/controller/rest/intent/create"

	appconfig "github.com/lelledaniele/upaygo/config"

	_ "github.com/lelledaniele/upaygo/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title uPayment in GO
// @version 1.0.0
// @description Microservice to manage payment
// @license.name MIT
func main() {
	s := appconfig.GetServerConfig()

	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(s.GetURI()+"/swagger/doc.json"),
	))
	http.HandleFunc(restintentcreate.URL, restintentcreate.Handler)

	log.Fatal(http.ListenAndServe(":"+s.GetPort(), nil))
}
