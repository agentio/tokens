package handlers

import (
	"log"
	"net/http"
	"slices"

	"github.com/agentio/atiquette/api/app/bsky"
	"github.com/agentio/tokens/internal/bureau"
	"github.com/agentio/tokens/internal/clients"
)

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	// When a user is signed-in, IO puts their did in this header.
	did := r.Header.Get("user-did")
	if did == "" {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}
	profile, err := bsky.ActorGetProfile(r.Context(), clients.AnonymousClient, did)
	if err != nil {
		log.Printf("%+v", err)
	}
	enrolled := slices.Contains(handles(), profile.Handle)
	if !enrolled {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	host := r.Host
	token, err := bureau.Issue(profile.Handle, profile.Did, *profile.DisplayName, host, 90)
	if err != nil {
		log.Printf("%+v", err)
	}
	response := token + "\n"
	w.Header().Set("content-disposition", `attachment; filename="token.txt"`)
	w.Write([]byte(response))
}
