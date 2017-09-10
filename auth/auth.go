package auth

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"salvation-army-api/model"
	"salvation-army-api/store"
	"time"
)

// AuthService provides authentication service
type AuthService interface {
	Auth(string, string) (map[string]interface{}, error)
}

type Authservice struct {
	Key     []byte
	Clients map[string]int
}

type CustomClaims struct {
	ClientID string `json:"clientId"`
	AccessLevel int `json:"access_level"`
	jwt.StandardClaims
}

const expiration = 12000000

func generateToken(signingKey []byte, clientID string) (string, error) {

	claims := CustomClaims{
		clientID,
		1,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * expiration).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func (as Authservice) Auth(clientID string, clientSecret string) (map[string]interface{}, error) {
	userStore := store.SqlUserStore{store.Database}
	me := <-userStore.GetByEmailAndPassword(clientID, clientSecret)
	if me.Err != nil {
		return map[string]interface{}{}, me.Err
	}
	signed, err := generateToken(as.Key, clientID)
	if err != nil {
		return map[string]interface{}{}, errors.New(err.Error())
	}
	return map[string]interface{}{"user":me.Data.(model.User),"token":signed}, nil
}

func MapClaimsFactory() jwt.Claims {
	return jwt.MapClaims{}
}

// ErrAuth is returned when credentials are incorrect
var ErrAuth = errors.New("Incorrect credentials")


