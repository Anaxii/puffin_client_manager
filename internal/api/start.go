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
		AllowedMethods: []string{"POST", "GET"},
	})

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/client/new", newClient).Methods("POST")
	r.HandleFunc("/client", getClient).Methods("GET")
	r.HandleFunc("/client/status", dAppStatus).Methods("GET")

	r.HandleFunc("/client/new/status", clientRequestStatus).Methods("POST")
	r.HandleFunc("/client/countries/update", updateCountries).Methods("POST")
	r.HandleFunc("/client/payments/next", paymentExpiration).Methods("GET")

	r.Use(mux.CORSMethodMiddleware(r))
	log.Info(fmt.Sprintf("API listening on port %v", port))
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(r)))
}
