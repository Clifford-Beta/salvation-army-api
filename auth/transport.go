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
		v, err := svc.Auth(req.ClientID, req.ClientSecret)
		if err != nil {
			return v, err
		}
		return v, nil
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type authRequest struct {
	ClientID     string    `json:"email"`
	ClientSecret string `json:"password"`
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
