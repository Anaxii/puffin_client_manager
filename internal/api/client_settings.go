package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"puffin_client_manager/pkg/client"
	"puffin_client_manager/pkg/db"
	"puffin_client_manager/pkg/global"
	"puffin_client_manager/pkg/util"
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

	id := util.RandStringRunes(20)
	client.QueueNewClient(requestBody, id)

	res, _ := json.Marshal(map[string]string{"status": "success", "id": id})
	w.Write(res)
	return
}

func clientRequestStatus(w http.ResponseWriter, r *http.Request) {
	type requestID struct {
		ID string `json:"id"`
	}
	var requestBody requestID
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status, err := db.Read([]byte("new_client"), []byte(requestBody.ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(map[string]string{"status": string(status)})
	w.Write(res)
	return
}

func updateCountry(w http.ResponseWriter, r *http.Request) {
	type clientSettings struct {
		Countries map[int][]string `json:"countries"`
		UUID      int              `json:"UUID"`
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
