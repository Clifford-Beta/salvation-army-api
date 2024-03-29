package school

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
	"salvation-army-api/model"
	"strconv"
	"time"
)

func MakeCreateEndpoint(svc SchoolService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.School)
		v, err := svc.Create(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeUpdateEndpoint(svc SchoolService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.School)
		v, err := svc.Update(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}


func MakeDeleteEndpoint(svc SchoolService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.School)
		v, err := svc.Delete(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRecordPerformanceEndpoint(svc SchoolService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.SchoolPerformance)
		v, err := svc.RecordPerformance(&req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRetrieveBestPerfomingSchoolEndpoint(svc SchoolService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schoolRequest)
		v, err := svc.GetBestSchool(req.From, req.To)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetOneEndpoint(svc SchoolService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schoolRequest)
		v, err := svc.GetOne(req.Id)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetAllEndpoint(svc SchoolService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(userRequest)
		v, err := svc.GetAll()
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRankSchoolByCategoryEndpoint(svc SchoolService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Query)
		v, err := svc.RankSchoolByCategory(req.Category,req.From,req.To)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeGetDashEndpoint(svc SchoolService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetDashboardData()
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRankAllSchoolsEndpoint(svc SchoolService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(schoolRequest)
		v, err := svc.RankAllSchools(req.From, req.To)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func DecodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.School
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRankByCatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request Query
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRecordPerformanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.SchoolPerformance
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetBestSchoolRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request schoolRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRankAllSchoolsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request schoolRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request schoolRequest
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return request, err
	}
	request.Id = id
	return request, nil
}

func DecodeGetAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request schoolRequest
	//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//	return nil, err
	//}
	return request, nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
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


type schoolRequest struct {
	Id   int `json:"id"`
	From int `json:"from"`
	To   int `json:"to"`
}

type Query struct {
	Category   int `json:"category"`
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

//type userResponse struct {
//	V   interface{} `json:"v"`
//	Err string `json:"err,omitempty"`
//}
