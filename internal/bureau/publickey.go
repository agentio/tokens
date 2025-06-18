package bureau

import (
	"os"

	"github.com/lestrrat-go/jwx/v3/jwk"
)

func PublicKey() (jwk.Key, error) {
	b, err := os.ReadFile("privatekey.json")
	if err != nil {
		return nil, err
	}
	key, err := jwk.ParseKey(b)
	if err != nil {
		return nil, err
	}
	return key.PublicKey()
}
