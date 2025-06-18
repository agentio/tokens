package check

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check JWT",
		Short: "check a JWT",
		RunE: func(cmd *cobra.Command, args []string) error {
			token := args[0]
			b, err := os.ReadFile("jwks-private.json")
			if err != nil {
				return err
			}
			privateKey, err := jwk.ParseKey(b)
			if err != nil {
				return err
			}
			publicKey, err := privateKey.PublicKey()
			if err != nil {
				return err
			}
			verifiedToken, err := jwt.Parse([]byte(token), jwt.WithKey(jwa.ES256(), publicKey))
			if err != nil {
				return err
			}
			b, err = json.MarshalIndent(verifiedToken, "", "  ")
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", string(b))
			return nil
		},
	}

	return cmd
}
