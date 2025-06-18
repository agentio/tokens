package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/agentio/tokens/pkg/bureau"
	"github.com/lestrrat-go/jwx/v3/jwk"
)

type JWKS struct {
	Keys []jwk.Key `json:"keys"`
}

func JWKSHandler(w http.ResponseWriter, r *http.Request) {
	key, err := bureau.PublicKey()
	if err != nil {
		return
	}
	b, err := json.Marshal(JWKS{
		Keys: []jwk.Key{key},
	})
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
