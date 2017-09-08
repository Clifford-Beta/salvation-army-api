package category

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
	"salvation-army-api/model"
	"strconv"
)

func MakeCreateEndpoint(svc CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.Category)
		v, err := svc.Create(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}
func MakeCreateTierEndpoint(svc CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.Tier)
		v, err := svc.CreateTier(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetOneEndpoint(svc CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(catRequest)
		v, err := svc.GetOne(req.Id)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetOneTierEndpoint(svc CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(catRequest)
		v, err := svc.GetOneTier(req.Id)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetAllEndpoint(svc CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(userRequest)
		v, err := svc.GetAll()
		if err != nil {
			return v, err
		}
		return v, nil
	}
}
func MakeGetAllTiersEndpoint(svc CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(userRequest)
		v, err := svc.GetAllTiers()
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func DecodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.Category
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func DecodeCreateTierRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.Tier
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request catRequest
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return request, err
	}
	request.Id = id
	return request, nil
}

func DecodeGetAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request catRequest
	return request, nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	//if err == nil {
	//	panic("encodeError with nil error")
	//}
	fmt.Println("Error", err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,GET,DELETE,PATCH")
	return json.NewEncoder(w).Encode(response)
}

type catRequest struct {
	Id int `json:"id"`
}

//type userResponse struct {
//	V   interface{} `json:"v"`
//	Err string `json:"err,omitempty"`
//}
