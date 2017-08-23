package activity

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

func MakeCreateActivityEndpoint(svc ActivityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.ExtraCurricular)
		v, err := svc.Create(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeCreateLevelEndpoint(svc ActivityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.ExtraCurricularLevel)
		v, err := svc.CreateLevel(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRecordPerformanceEndpoint(svc ActivityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.ExtraCurricularActivity)
		v, err := svc.RecordPerformance(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetOneActivityEndpoint(svc ActivityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(catRequest)
		v, err := svc.GetOneActivity(req.Id)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetOneLevelEndpoint(svc ActivityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(catRequest)
		v, err := svc.GetOneLevel(req.Id)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetOnePerformanceEndpoint(svc ActivityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(catRequest)
		v, err := svc.GetOnePerformance(req.Id)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetAllActivitiesEndpoint(svc ActivityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(userRequest)
		v, err := svc.GetAllActivities()
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetAllLevelsEndpoint(svc ActivityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(userRequest)
		v, err := svc.GetAllLevels()
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetAllPerformancesEndpoint(svc ActivityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(userRequest)
		v, err := svc.GetAllPerformances()
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func DecodeCreateActivityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.ExtraCurricular
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeCreateLevelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.ExtraCurricularLevel
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRecordPerformanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.ExtraCurricularActivity
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
	return json.NewEncoder(w).Encode(response)
}

type catRequest struct {
	Id int `json:"id"`
}

//type userResponse struct {
//	V   interface{} `json:"v"`
//	Err string `json:"err,omitempty"`
//}
