package infrastructure

import (
	"salvation-army-api/model"
	"salvation-army-api/store"
	"time"
)

// StringService provides operations on strings.
type InfrastructureService interface {
	Create(message model.Infrastructure) (*model.Infrastructure, error)
	CreateType(message model.InfrastructureType) (*model.InfrastructureType, error)
	GetOne(int) (model.Infrastructure, error)
	GetOneType(int) (model.InfrastructureType, error)
	GetAll() (map[string][]*model.Infrastructure, error)
	GetAllTypes() (map[string][]*model.InfrastructureType, error)
}

type Infrastructureservice struct{}

func (Infrastructureservice) Create(inf model.Infrastructure) (*model.Infrastructure, error) {
	iStore := store.SqlInfrastructureStore{store.Database}
	inf.Status = 1
	inf.DateCreated = time.Now()
	me := <-iStore.Create(&inf)
	if me.Err != nil {
		return &model.Infrastructure{}, me.Err
	}
	return me.Data.(*model.Infrastructure), nil
}

func (Infrastructureservice) CreateType(inf model.InfrastructureType) (*model.InfrastructureType, error) {
	iStore := store.SqlInfrastructureStore{store.Database}
	inf.Status = 1
	me := <-iStore.CreateIType(&inf)
	if me.Err != nil {
		return &model.InfrastructureType{}, me.Err
	}
	return me.Data.(*model.InfrastructureType), nil
}

func (Infrastructureservice) GetOne(id int) (model.Infrastructure, error) {
	iStore := store.SqlInfrastructureStore{store.Database}
	me := <-iStore.RetrieveOne(id)
	if me.Err != nil {
		return model.Infrastructure{}, me.Err
	}
	return me.Data.(model.Infrastructure), nil
}
func (Infrastructureservice) GetOneType(id int) (model.InfrastructureType, error) {
	iStore := store.SqlInfrastructureStore{store.Database}
	me := <-iStore.RetrieveOneType(id)
	if me.Err != nil {
		return model.InfrastructureType{}, me.Err
	}
	return me.Data.(model.InfrastructureType), nil
}

func (Infrastructureservice) GetAll() (map[string][]*model.Infrastructure, error) {
	iStore := store.SqlInfrastructureStore{store.Database}
	me := <-iStore.RetrieveAll()
	if me.Err != nil {
		return map[string][]*model.Infrastructure{"data": []*model.Infrastructure{}}, me.Err
	}
	return map[string][]*model.Infrastructure{"data": me.Data.([]*model.Infrastructure)}, nil
}

func (Infrastructureservice) GetAllTypes() (map[string][]*model.InfrastructureType, error) {
	iStore := store.SqlInfrastructureStore{store.Database}
	me := <-iStore.RetrieveAllTypes()
	if me.Err != nil {
		return map[string][]*model.InfrastructureType{"data": []*model.InfrastructureType{}}, me.Err
	}
	return map[string][]*model.InfrastructureType{"data": me.Data.([]*model.InfrastructureType)}, nil
}

//func (Userservice) Update(a,b int) (int, error) {
//	return a+b, nil
//}
