package school

import (
	"salv_prj/model"
	"salv_prj/store"
	"time"
)

type SchoolService interface {
	Create(model.School)(*model.School, error)
	GetOne(int)(model.School,error)
	GetAll()([]*model.School,error)
	RecordPerformance(performance *model.SchoolPerformance)(*model.SchoolPerformance,error)
	GetBestSchool(from,to int)(model.SchoolPerformanceResult,error)
	RankAllSchools(from,to int)([]*model.SchoolPerformanceResult,error)
}

type Schoolservice struct {}

func ( Schoolservice)Create(school model.School)(*model.School,error)  {
	schoolStore := store.SqlSchoolStore{store.Database}
	school.Status = 1
	school.DateRegistered = time.Now()
	sch := <- schoolStore.Save(&school)
	if sch.Err != nil {
		return &model.School{},sch.Err
	}
	return sch.Data.(*model.School),nil

}

func (Schoolservice)GetOne(id int)(model.School,error)  {
	schoolStore := store.SqlSchoolStore{store.Database}
	sch := <- schoolStore.Get(id)
	if sch.Err != nil {
		return model.School{},sch.Err
	}
	return sch.Data.(model.School),nil

}

func (Schoolservice) GetAll()([]*model.School,error) {
	schoolStore := store.SqlSchoolStore{store.Database}
	sch := <- schoolStore.GetMany()
	if sch.Err != nil {
		return []*model.School{},sch.Err
	}
	return sch.Data.([]*model.School),nil

}

func (Schoolservice) RecordPerformance(performance *model.SchoolPerformance)(*model.SchoolPerformance,error) {
	schoolStore := store.SqlSchoolStore{store.Database}
	pf := <- schoolStore.RecordPerformance(performance)
	if pf.Err != nil {
		return &model.SchoolPerformance{},pf.Err
	}
	return pf.Data.(*model.SchoolPerformance),nil
}

func (Schoolservice) GetBestSchool(from,to int)(model.SchoolPerformanceResult,error) {
	schoolStore := store.SqlSchoolStore{store.Database}
	filters := map[string]interface{}{}
	if from != 0 {
		filters["from"] = from
		if to != 0 {
			filters["to"] = to

		}else{
			filters["to"] = from

		}
	}
	bestSchool := <- schoolStore.RetrieveBestPerfomingSchool(filters)
	if bestSchool.Err != nil {
		return model.SchoolPerformanceResult{},bestSchool.Err
	}
	return bestSchool.Data.(model.SchoolPerformanceResult),nil
}

func (Schoolservice) RankAllSchools(from,to int)([]*model.SchoolPerformanceResult,error) {
	schoolStore := store.SqlSchoolStore{store.Database}
	filters := map[string]interface{}{}
	if from != 0 {
		filters["from"] = from
		if to != 0 {
			filters["to"] = to

		}else{
			filters["to"] = from

		}
	}
	bestSchool := <- schoolStore.RankAllSchools()
	if bestSchool.Err != nil {
		return []*model.SchoolPerformanceResult{},bestSchool.Err
	}
	return bestSchool.Data.([]*model.SchoolPerformanceResult),nil
}