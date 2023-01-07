package api

import (
	"encoding/json"
	"net/http"
)

func addCountry(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(map[string]string{"error": "country already blocked"})
	w.Write(res)
	return
}

func removeCountry(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(map[string]string{"error": "country already blocked"})
	w.Write(res)
	return
}

func countries(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(map[string]string{"error": "country already blocked"})
	w.Write(res)
	return
}

func paymentExpiration(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(map[string]string{"error": "country already blocked"})
	w.Write(res)
	return
}