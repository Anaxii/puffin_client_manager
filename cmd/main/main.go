package main

import (
	"puffin_client_manager/internal/api"
	"puffin_client_manager/internal/config"
	"puffin_client_manager/internal/database"
	"puffin_client_manager/internal/payments"
	"puffin_client_manager/pkg/client"
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

	go paymentsHandler.StartPaymentsHandler()
	go client.ListenForNewClients()
	api.StartAPI("8080")
}
