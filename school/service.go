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
}

type Schoolservice struct {}

func (Schoolservice)Create(school model.School)(*model.School,error)  {
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