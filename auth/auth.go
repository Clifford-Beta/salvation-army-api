package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// AuthService provides authentication service
type AuthService interface {
	Auth(string, string) (string, error)
}

type Authservice struct {
	Key     []byte
	Clients map[string]string
}

type CustomClaims struct {
	ClientID string `json:"clientId"`
	jwt.StandardClaims
}

const expiration = 1200

func generateToken(signingKey []byte, clientID string) (string, error) {
	claims := CustomClaims{
		clientID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * expiration).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func (as Authservice) Auth(clientID string, clientSecret string) (string, error) {

	if as.Clients[clientID] == clientSecret {
		signed, err := generateToken(as.Key, clientID)
		if err != nil {
			return "", errors.New(err.Error())
		}
		return signed, nil
	}
	return "", ErrAuth
}

// ErrAuth is returned when credentials are incorrect
var ErrAuth = errors.New("Incorrect credentials")