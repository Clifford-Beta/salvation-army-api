package activity

import (
	"salvation-army-api/model"
	"salvation-army-api/store"
	"time"
)

type ActivityService interface {
	Create(activity model.ExtraCurricular) (*model.ExtraCurricular, error)
	Update(activity model.ExtraCurricular) (bool, error)
	CreateLevel(level model.ExtraCurricularLevel) (*model.ExtraCurricularLevel, error)
	RecordPerformance(performance model.ExtraCurricularActivity) (*model.ExtraCurricularActivity, error)
	GetOneActivity(id int) (model.ExtraCurricular, error)
	GetOneLevel(id int) (model.ExtraCurricularLevel, error)
	GetOnePerformance(id int) (model.ExtraCurricularActivity, error)
	GetAllActivities() (map[string][]*model.ExtraCurricular, error)
	GetAllLevels() (map[string][]*model.ExtraCurricularLevel, error)
	GetAllPerformances() (map[string][]*model.ExtraCurricularActivityResult, error)
}
type Activityservice struct{}

func (Activityservice) Create(act model.ExtraCurricular) (*model.ExtraCurricular, error) {
	actStore := store.SqlExtraCurricularStore{store.Database}
	act.Status = 1
	act.TimeStamp = time.Now()
	if err := act.Validate(); err != nil {
		return &model.ExtraCurricular{},err
	}
	res := <-actStore.CreateNewActivity(&act)
	if res.Err != nil {
		return &model.ExtraCurricular{}, res.Err
	}
	return res.Data.(*model.ExtraCurricular), nil
}

func (Activityservice) Update(msg model.ExtraCurricular) (bool, error) {
	projStore := store.SqlExtraCurricularStore{store.Database}
	me := <-projStore.Update(&msg)
	if me.Err != nil {
		return me.Data.(bool), me.Err
	}
	return me.Data.(bool), nil
}


func (Activityservice) CreateLevel(act model.ExtraCurricularLevel) (*model.ExtraCurricularLevel, error) {
	actStore := store.SqlExtraCurricularStore{store.Database}
	act.Status = 1
	act.TimeStamp = time.Now()
	if err := act.Validate(); err != nil {
		return &model.ExtraCurricularLevel{},err
	}
	res := <-actStore.RecordLevel(&act)
	if res.Err != nil {
		return &model.ExtraCurricularLevel{}, res.Err
	}
	return res.Data.(*model.ExtraCurricularLevel), nil
}

func (Activityservice) RecordPerformance(act model.ExtraCurricularActivity) (*model.ExtraCurricularActivity, error) {
	actStore := store.SqlExtraCurricularStore{store.Database}
	act.Status = 1
	act.TimeStamp = time.Now()
	if err := act.Validate(); err != nil {
		return &model.ExtraCurricularActivity{},err
	}
	res := <-actStore.RecordActivity(&act)
	if res.Err != nil {
		return &model.ExtraCurricularActivity{}, res.Err
	}
	return res.Data.(*model.ExtraCurricularActivity), nil
}

func (Activityservice) GetOneActivity(id int) (model.ExtraCurricular, error) {
	actStore := store.SqlExtraCurricularStore{store.Database}
	sch := <-actStore.GetActivity(id)
	if sch.Err != nil {
		return model.ExtraCurricular{}, sch.Err
	}
	return sch.Data.(model.ExtraCurricular), nil

}

func (Activityservice) GetOneLevel(id int) (model.ExtraCurricularLevel, error) {
	actStore := store.SqlExtraCurricularStore{store.Database}
	sch := <-actStore.GetLevel(id)
	if sch.Err != nil {
		return model.ExtraCurricularLevel{}, sch.Err
	}
	return sch.Data.(model.ExtraCurricularLevel), nil

}

func (Activityservice) GetOnePerformance(id int) (model.ExtraCurricularActivity, error) {
	actStore := store.SqlExtraCurricularStore{store.Database}
	sch := <-actStore.GetRecordedActivity(id)
	if sch.Err != nil {
		return model.ExtraCurricularActivity{}, sch.Err
	}
	return sch.Data.(model.ExtraCurricularActivity), nil

}

func (Activityservice) GetAllActivities() (map[string][]*model.ExtraCurricular, error) {
	actStore := store.SqlExtraCurricularStore{store.Database}
	sch := <-actStore.GetAllActivities()
	if sch.Err != nil {
		return map[string][]*model.ExtraCurricular{"data": []*model.ExtraCurricular{}}, sch.Err
	}
	return map[string][]*model.ExtraCurricular{"data": sch.Data.([]*model.ExtraCurricular)}, nil

}

func (Activityservice) GetAllLevels() (map[string][]*model.ExtraCurricularLevel, error) {
	actStore := store.SqlExtraCurricularStore{store.Database}
	sch := <-actStore.GetAllLevels()
	if sch.Err != nil {
		return map[string][]*model.ExtraCurricularLevel{"data": []*model.ExtraCurricularLevel{}}, sch.Err
	}
	return map[string][]*model.ExtraCurricularLevel{"data": sch.Data.([]*model.ExtraCurricularLevel)}, nil

}

func (Activityservice) GetAllPerformances() (map[string][]*model.ExtraCurricularActivityResult, error) {
	actStore := store.SqlExtraCurricularStore{store.Database}
	sch := <-actStore.GetAllRecordedActivities()
	if sch.Err != nil {
		return map[string][]*model.ExtraCurricularActivityResult{"data": []*model.ExtraCurricularActivityResult{}}, sch.Err
	}
	return map[string][]*model.ExtraCurricularActivityResult{"data": sch.Data.([]*model.ExtraCurricularActivityResult)}, nil
}
