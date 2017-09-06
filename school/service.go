package school

import (
	"salvation-army-api/model"
	"salvation-army-api/store"
	"time"
)

type SchoolService interface {
	Create(model.School) (*model.School, error)
	Update(model.School) (bool, error)
	GetOne(int) (model.School, error)
	GetAll() (map[string][]*model.SchoolResult, error)
	RecordPerformance(performance *model.SchoolPerformance) (*model.SchoolPerformance, error)
	//UpdatePerformance(performance *model.SchoolPerformance) (*model.SchoolPerformance, error)
	//RetrievePerformance(int) (model.SchoolPerformance, error)
	GetBestSchool(from, to int) (model.SchoolPerformanceResult, error)
	RankAllSchools(from, to int) (map[string][]model.SchoolPerformanceResult, error)
}

type Schoolservice struct{}

func (Schoolservice) Create(school model.School) (*model.School, error) {
	schoolStore := store.SqlSchoolStore{store.Database}
	school.Status = 1
	school.DateRegistered = time.Now()
	if err := school.Validate(); err != nil {
		return &model.School{},err
	}
	sch := <-schoolStore.Save(&school)
	if sch.Err != nil {
		return &model.School{}, sch.Err
	}
	return sch.Data.(*model.School), nil

}

func (Schoolservice) GetOne(id int) (model.School, error) {
	schoolStore := store.SqlSchoolStore{store.Database}
	sch := <-schoolStore.Get(id)
	if sch.Err != nil {
		return model.School{}, sch.Err
	}
	return sch.Data.(model.School), nil

}

func (Schoolservice) Update(school model.School) (bool, error) {
	userStore := store.SqlSchoolStore{store.Database}
	me := <-userStore.Update(&school)
	if me.Err != nil {
		return me.Data.(bool), me.Err
	}
	return me.Data.(bool), nil
}


func (Schoolservice) GetAll() (map[string][]*model.SchoolResult, error) {
	schoolStore := store.SqlSchoolStore{store.Database}
	sch := <-schoolStore.GetMany()
	if sch.Err != nil {
		return map[string][]*model.SchoolResult{"data": []*model.SchoolResult{}}, sch.Err
	}
	return map[string][]*model.SchoolResult{"data": sch.Data.([]*model.SchoolResult)}, nil

}

func (Schoolservice) RecordPerformance(performance *model.SchoolPerformance) (*model.SchoolPerformance, error) {
	schoolStore := store.SqlSchoolStore{store.Database}
	pf := <-schoolStore.RecordPerformance(performance)
	if pf.Err != nil {
		return &model.SchoolPerformance{}, pf.Err
	}
	return pf.Data.(*model.SchoolPerformance), nil
}

func (Schoolservice) GetBestSchool(from, to int) (model.SchoolPerformanceResult, error) {
	schoolStore := store.SqlSchoolStore{store.Database}
	filters := map[string]interface{}{}
	if from != 0 {
		filters["from"] = from
		if to != 0 {
			filters["to"] = to

		} else {
			filters["to"] = from

		}
	}
	bestSchool := <-schoolStore.RetrieveBestPerfomingSchool(filters)
	if bestSchool.Err != nil {
		return model.SchoolPerformanceResult{}, bestSchool.Err
	}
	return bestSchool.Data.(model.SchoolPerformanceResult), nil
}

func (Schoolservice) RankAllSchools(from, to int) (map[string][]model.SchoolPerformanceResult, error) {
	schoolStore := store.SqlSchoolStore{store.Database}
	filters := map[string]interface{}{}
	if from != 0 {
		filters["from"] = from
		if to != 0 {
			filters["to"] = to

		} else {
			filters["to"] = from

		}
	}
	bestSchool := <-schoolStore.RankAllSchools()
	if bestSchool.Err != nil {
		return map[string][]model.SchoolPerformanceResult{"data": []model.SchoolPerformanceResult{}}, bestSchool.Err
	}
	return map[string][]model.SchoolPerformanceResult{"data": bestSchool.Data.([]model.SchoolPerformanceResult)}, nil
}
