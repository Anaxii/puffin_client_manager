package main

import (
	"puffin_client_manager/internal/config"
	"puffin_client_manager/internal/database"
	"puffin_client_manager/internal/payments"
	Log "puffin_client_manager/pkg/log"
)

func main() {
	Log.SetupLogs()

	c := config.GetConfig()

	paymentsHandler := payments.PaymentsHandler{
		Config: c,
		DB: database.Database{
			DBURI: c.MongoDbURI,
		},
	}
	paymentsHandler.StartPaymentsHandler()
}
