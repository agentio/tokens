package bureau

import (
	"errors"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

// Issue a JWT using a local private key
func Issue(handle, did, name, host string, lifetime int) (string, error) {
	b, err := os.ReadFile("privatekey.json")
	if err != nil {
		return "", err
	}
	key, err := jwk.ParseKey(b)
	if err != nil {
		return "", err
	}
	if handle == "" || name == "" {
		return "", errors.New("handle and name must be specified")
	}
	var payload []byte
	token := jwt.New()
	token.Set(jwt.SubjectKey, handle)
	token.Set(jwt.AudienceKey, `application`)
	token.Set(jwt.IssuerKey, "https://"+host)
	token.Set(jwt.IssuedAtKey, time.Now())
	token.Set("name", name)
	if did != "" {
		token.Set("did", did)
	}
	// tokens expire in 24 hours
	exp := time.Now().Add(24 * time.Hour)
	token.Set(jwt.ExpirationKey, exp)
	payload, err = jwt.Sign(token, jwt.WithKey(jwa.ES256(), key))
	if err != nil {
		return "", err
	}
	return string(payload), nil
}
