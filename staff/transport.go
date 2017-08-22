package staff

import (
	"context"
	"encoding/json"
	"net/http"
	"salv_prj/model"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

func MakeAddStaffEndpoint(svc StaffService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.Staff)
		v, err := svc.AddStaff(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}
func MakeRetrieveStaffEndpoint(svc StaffService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(staffRequest)
		v, err := svc.RetrieveStaff(req.Id)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRetrieveAllStaffEndpoint(svc StaffService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.RetrieveAllStaff()
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeAddStaffRoleEndpoint(svc StaffService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.StaffRole)
		v, err := svc.AddStaffRole(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRetrieveStaffRoleEndpoint(svc StaffService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(staffRequest)
		v, err := svc.RetrieveStaffRole(req.Id)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRecordBestPerformingStaffEndpoint(svc StaffService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BestTeacher)
		v, err := svc.RecordBestPerformingStaff(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRecordBestPerformingStudentEndpoint(svc StaffService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BestStudent)
		v, err := svc.RecordBestPerformingStudent(req)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRetrieveBestPerformingStaffEndpoint(svc StaffService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(staffRequest)
		v, err := svc.RetrieveBestPerformingStaff(req.From,req.To)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRetrieveBestPerformingStudentEndpoint(svc StaffService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(staffRequest)
		v, err := svc.RetrieveBestPerformingStudent(req.From,req.To)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRankStaffPerformanceEndpoint(svc StaffService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(staffRequest)
		v, err := svc.RankStaffPerformance(req.From,req.To)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func MakeRankStudentPerformanceEndpoint(svc StaffService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(staffRequest)
		v, err := svc.RankStudentPerformance(req.From,req.To)
		if err != nil {
			return v, err
		}
		return v, nil
	}
}

func DecodeAddStaffRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.Staff
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeAddStaffRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.StaffRole
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRecordBestPerformingStaffRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.BestTeacher
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRecordBestPerformingStudentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.BestStudent
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}


func DecodeRetrieveStaffRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request staffRequest
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if  err != nil {
		return request,err
	}
	request.Id = id
	return request, nil
}

func DecodeRetrieveStaffRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request staffRequest
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if  err != nil {
		return request,err
	}
	request.Id = id
	return request, nil
}

func DecodeRetrieveAllStaffRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request staffRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRetrieveBestPerformingStaffRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request staffRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRetrieveBestPerformingStudentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request staffRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRankStaffPerformanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request staffRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRankStudentPerformanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request staffRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeRetrieveAllRolesStaffRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request staffRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
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
type staffRequest struct {
	Id int `json:"id"`
	From int `json:"from"`
	To int `json:"to"`
}

//type userResponse struct {
//	V   interface{} `json:"v"`
//	Err string `json:"err,omitempty"`
//}
