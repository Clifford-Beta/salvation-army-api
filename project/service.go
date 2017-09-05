package project

import (
	"salvation-army-api/model"
	"salvation-army-api/store"
	"time"
)

// StringService provides operations on strings.
type ProjectService interface {
	Create(project model.Project) (*model.Project, error)
	GetOne(int) (model.Project, error)
	GetAll() (map[string][]model.ProjectResult, error)
	Update(project model.Project)(bool,error)
}

type Projectservice struct{}

func (Projectservice) Create(project model.Project) (*model.Project, error) {
	projStore := store.SqlProjectStore{store.Database}
	project.Status = 1
	project.Start = time.Now()
	project.TimeStamp = time.Now()
	if err := project.Validate(); err != nil {
		return &model.Project{},err
	}
	me := <-projStore.Create(&project)
	if me.Err != nil {
		return &model.Project{}, me.Err
	}
	return me.Data.(*model.Project), nil
}

func (Projectservice) Update(project model.Project) (bool, error) {
	projStore := store.SqlProjectStore{store.Database}
	me := <-projStore.Update(&project)
	if me.Err != nil {
		return me.Data.(bool), me.Err
	}
	return me.Data.(bool), nil
}

func (Projectservice) GetOne(id int) (model.Project, error) {
	projStore := store.SqlProjectStore{store.Database}
	me := <-projStore.Retrieve(id)
	if me.Err != nil {
		return model.Project{}, me.Err
	}
	return me.Data.(model.Project), nil
}

func (Projectservice) GetAll() (map[string][]model.ProjectResult, error) {
	projStore := store.SqlProjectStore{store.Database}
	me := <-projStore.RetrieveAll()
	if me.Err != nil {
		return map[string][]model.ProjectResult{"data": []model.ProjectResult{}}, me.Err
	}
	return map[string][]model.ProjectResult{"data": me.Data.([]model.ProjectResult)}, nil
}
