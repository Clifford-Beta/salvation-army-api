package infrastructure

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

func MakeCreateEndpoint(svc InfrastructureService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.Infrastructure)
		v, err := svc.Create(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeUpdateEndpoint(svc InfrastructureService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.Infrastructure)
		v, err := svc.Update(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeDeleteEndpoint(svc InfrastructureService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.Infrastructure)
		v, err := svc.Delete(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeCreateTypeEndpoint(svc InfrastructureService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.InfrastructureType)
		v, err := svc.CreateType(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetOneEndpoint(svc InfrastructureService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(msgRequest)
		v, err := svc.GetOne(req.Id)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetOneTypeEndpoint(svc InfrastructureService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(msgRequest)
		v, err := svc.GetOneType(req.Id)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetAllEndpoint(svc InfrastructureService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetAll()
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetAllTypesEndpoint(svc InfrastructureService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetAllTypes()
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func DecodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.Infrastructure
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeCreateTypeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.InfrastructureType
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request msgRequest
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return request, err
	}
	request.Id = id
	return request, nil
}

func DecodeGetAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request msgRequest
	//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//	return nil, err
	//}
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

type msgRequest struct {
	Id int `json:"id"`
}

//type userResponse struct {
//	V   interface{} `json:"v"`
//	Err string `json:"err,omitempty"`
//}
