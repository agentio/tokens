package handlers

import (
	"encoding/json"
	"os"
)

func handles() []string {
	b, err := os.ReadFile("users.json")
	if err != nil {
		return []string{}
	}
	var h []string
	err = json.Unmarshal(b, &h)
	if err != nil {
		return []string{}
	}
	return h
}
