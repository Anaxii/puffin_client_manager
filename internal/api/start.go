package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func StartAPI(port string) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST"},
	})

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/newClient", newClient).Methods("POST")
	r.HandleFunc("/updateCountry", updateCountry).Methods("POST")
	r.HandleFunc("/paymentExpiration", paymentExpiration).Methods("GET")

	r.Use(mux.CORSMethodMiddleware(r))
	log.Info(fmt.Sprintf("API listening on port %v", port))
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(r)))
}
