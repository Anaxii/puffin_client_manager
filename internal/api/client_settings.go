package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"puffin_client_manager/pkg/client"
	"puffin_client_manager/pkg/global"
)

func newClient(w http.ResponseWriter, r *http.Request) {
	var requestBody global.ClientSettings
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.Warn("Failed to decode request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client.QueueNewClient(requestBody)

	res, _ := json.Marshal(map[string]string{"status": "success"})
	w.Write(res)
	return
}

func updateCountry(w http.ResponseWriter, r *http.Request) {
	type clientSettings struct {
		Countries map[int][]string `json:"countries"`
		UUID      int            `json:"UUID"`
	}

	res, _ := json.Marshal(map[string]string{"error": "country already added"})
	w.Write(res)
	return
}


func paymentExpiration(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(map[string]string{"error": "country already blocked"})
	w.Write(res)
	return
}