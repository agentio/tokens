package serve

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agentio/tokens/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// This app listens on this port.
const port = 3000

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "run the token server",
		RunE: func(cmd *cobra.Command, args []string) error {
			mux := mux.NewRouter()
			mux.HandleFunc("/", handlers.HomeHandler)
			mux.HandleFunc("/signin", handlers.SigninHandler)
			mux.HandleFunc("/generate", handlers.GenerateHandler)
			mux.HandleFunc("/jwks.json", handlers.JWKSHandler)
			mux.HandleFunc("/.well-known/openid-configuration", handlers.OpenIDConfigurationHandler)
			mux.HandleFunc("/public/styles.css", handlers.StylesHandler)
			mux.HandleFunc("/public/icon.svg", handlers.IconHandler)
			httpd := &http.Server{
				Addr:    fmt.Sprintf(":%d", port),
				Handler: mux,
			}
			if err := httpd.ListenAndServe(); err != nil {
				log.Fatalf("%s", err)
			}
			return nil
		},
	}

	return cmd
}
