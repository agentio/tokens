package generate

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "generate a JWT",
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := os.ReadFile("privatekey.json")
			if err != nil {
				return err
			}
			privKey, err := jwk.ParseKey(b)
			if err != nil {
				return err
			}
			var payload []byte
			{ // Create signed payload
				token := jwt.New()
				token.Set(jwt.SubjectKey, `Tim Burks`)
				token.Set(jwt.AudienceKey, `io`)
				token.Set(jwt.IssuerKey, "https://agent.io")
				token.Set(jwt.IssuedAtKey, time.Now())

				exp, err := time.Parse(time.DateOnly, "2026-01-01")
				if err != nil {
					log.Fatalf("%s", err)
				}
				token.Set(jwt.ExpirationKey, exp)
				token.Set(jwt.JwtIDKey, "1")
				//SubjectKey    = "sub"
				//AudienceKey   = "aud"
				//IssuerKey     = "iss"
				//IssuedAtKey   = "iat"
				//ExpirationKey = "exp"
				//JwtIDKey      = "jti"
				//NotBeforeKey  = "nbf"

				payload, err = jwt.Sign(token, jwt.WithKey(jwa.ES256(), privKey))
				if err != nil {
					return err
				}
			}
			fmt.Printf("%s\n", string(payload))

			return nil
		},
	}

	return cmd
}
