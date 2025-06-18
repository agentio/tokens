package handlers

import (
	"encoding/json"
	"net/http"
)

func OpenIDConfigurationHandler(w http.ResponseWriter, r *http.Request) {
	config := &OpenIDConfiguration{
		JwksUri: "https://" + r.Host + "/jwks.json",
	}
	b, err := json.Marshal(config)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

type OpenIDConfiguration struct {
	JwksUri string `json:"jwks_uri"`
}
