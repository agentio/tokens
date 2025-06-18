package generate

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "generate a json representation of a private key",
		RunE: func(cmd *cobra.Command, args []string) error {
			privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
			if err != nil {
				return err
			}
			key, err := jwk.Import(privKey)
			if err != nil {
				return err
			}
			// use the current time as the key id
			kid := fmt.Sprintf("%d", time.Now().Unix())
			if err := key.Set(jwk.KeyIDKey, kid); err != nil {
				return err
			}
			b, err := json.MarshalIndent(key, "", "  ")
			if err != nil {
				return err
			}
			if err := os.WriteFile("privatekey.json", b, 0644); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
