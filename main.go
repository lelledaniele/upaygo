package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	cic "github.com/lelledaniele/upaygo/controller/api/intent/create"

	conf "github.com/lelledaniele/upaygo/config"

	_ "github.com/lelledaniele/upaygo/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title uPayment in GO
// @version 1.0.0
// @description Microservice to manage payment
// @license.name MIT
func main() {
	u, p := conf.GetServerConfig().GetURI(), conf.GetServerConfig().GetPort()
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(u+"/swagger/doc.json"),
	))

	_ = http.ListenAndServe(":"+p, r)

	http.HandleFunc(cic.URL, cic.Handler)

	log.Fatal(http.ListenAndServe(":"+p, nil))
}
