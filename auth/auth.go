package auth

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"salv_prj/model"
	"salv_prj/store"
	"time"
)

// AuthService provides authentication service
type AuthService interface {
	Auth(int, string) (string, error)
}

type Authservice struct {
	Key     []byte
	Clients map[string]int
}

type CustomClaims struct {
	ClientID string `json:"clientId"`
	jwt.StandardClaims
}

const expiration = 1200000000

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

func (as Authservice) Auth(clientID int, clientSecret string) (string, error) {
	clientStore := store.SqlUserStore{store.Database}
	client := <-clientStore.Get(clientID)
	if client.Err != nil {
		return "", ErrAuth
	}
	log.Println("this is the user", client.Data)
	signed, err := generateToken(as.Key, client.Data.(model.User).Name)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return signed, nil
	//if as.Clients[clientID] == clientSecret {
	//	signed, err := generateToken(as.Key, client.Data.(model.User))
	//	if err != nil {
	//		return "", errors.New(err.Error())
	//	}
	//	return signed, nil
	//}
}

// ErrAuth is returned when credentials are incorrect
var ErrAuth = errors.New("Incorrect credentials")
