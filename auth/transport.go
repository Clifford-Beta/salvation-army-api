package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeAuthEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(authRequest)
		token, err := svc.Auth(req.ClientID, req.ClientSecret)
		if err != nil {
			return nil, err
		}
		return authResponse{token, ""}, nil
	}
}

func DecodeAuthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request authRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return json.NewEncoder(w).Encode(response)
}

type authRequest struct {
	ClientID     int    `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type authResponse struct {
	Token string `json:"token,omitempty"`
	Err   string `json:"err,omitempty"`
}

func AuthErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	code := http.StatusUnauthorized
	msg := err.Error()

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(authResponse{Token: "", Err: msg})
}
