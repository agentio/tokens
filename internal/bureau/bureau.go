package bureau

import (
	"errors"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
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
	{ // Create signed payload
		token := jwt.New()
		token.Set(jwt.SubjectKey, handle)
		token.Set(jwt.AudienceKey, `application`)
		token.Set(jwt.IssuerKey, "https://"+host)
		token.Set(jwt.IssuedAtKey, time.Now())
		if did != "" {
			token.Set("did", did)
		}
		if name != "" {
			token.Set("name", name)
		}
		/*
			// one idea for expiration
			exp := time.Now().Add(1 * time.Hour * 24 * time.Duration(lifetime))
			year, month, day := exp.Date()
		*/
		/*
			// another idea for expiration
			exp := time.Unix(9999999999, 0)
		*/
		year := 2199
		month := time.Month(12)
		day := 31
		exp := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		token.Set(jwt.ExpirationKey, exp)
		// I don't think the "jti" key is neededd
		if false {
			id := exp.Format("20060102") + "-" + handle
			token.Set(jwt.JwtIDKey, id)
		}
		payload, err = jwt.Sign(token, jwt.WithKey(jwa.ES256(), key))
		if err != nil {
			return "", err
		}
	}
	return string(payload), nil
}
