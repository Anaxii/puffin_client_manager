package main

import (
	log "github.com/sirupsen/logrus"
	"puffin_client_manager/internal/config"
	Log "puffin_client_manager/pkg/log"
)

func main() {
	Log.SetupLogs()

	c := config.GetConfig()
	log.Println(c)
}
