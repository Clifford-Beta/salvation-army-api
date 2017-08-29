package category

import (
	"salvation-army-api/model"
	"salvation-army-api/store"
)

type CategoryService interface {
	Create(category model.Category) (*model.Category, error)
	CreateTier(category model.Tier) (*model.Tier, error)
	GetOne(id int) (model.Category, error)
	GetOneTier(id int) (model.Tier, error)
	GetAll() (map[string][]*model.Category, error)
	GetAllTiers() (map[string][]*model.Tier, error)
}
type Categoryservice struct{}

func (Categoryservice) Create(category model.Category) (*model.Category, error) {
	catStore := store.SqlCategoryStore{store.Database}
	category.Status = 1
	if err := category.Validate(); err != nil {
		return &model.Category{},err
	}
	res := <-catStore.Save(&category)
	if res.Err != nil {
		return &model.Category{}, res.Err
	}
	return res.Data.(*model.Category), nil
}
func (Categoryservice) CreateTier(category model.Tier) (*model.Tier, error) {
	catStore := store.SqlCategoryStore{store.Database}
	if err := category.Validate(); err != nil {
		return &model.Tier{},err
	}
	res := <-catStore.CreateTier(&category)
	if res.Err != nil {
		return &model.Tier{}, res.Err
	}
	return res.Data.(*model.Tier), nil
}

func (Categoryservice) GetOne(id int) (model.Category, error) {
	catStore := store.SqlCategoryStore{store.Database}
	sch := <-catStore.Get(id)
	if sch.Err != nil {
		return model.Category{}, sch.Err
	}
	return sch.Data.(model.Category), nil

}
func (Categoryservice) GetOneTier(id int) (model.Tier, error) {
	catStore := store.SqlCategoryStore{store.Database}
	sch := <-catStore.GetTier(id)
	if sch.Err != nil {
		return model.Tier{}, sch.Err
	}
	return sch.Data.(model.Tier), nil

}

func (Categoryservice) GetAll() (map[string][]*model.Category, error) {
	catStore := store.SqlCategoryStore{store.Database}
	sch := <-catStore.GetMany()
	if sch.Err != nil {
		return map[string][]*model.Category{}, sch.Err
	}
	return map[string][]*model.Category{"data":sch.Data.([]*model.Category)}, nil

}

func (Categoryservice) GetAllTiers() (map[string][]*model.Tier, error) {
	catStore := store.SqlCategoryStore{store.Database}
	sch := <-catStore.GetManyTiers()
	if sch.Err != nil {
		return map[string][]*model.Tier{"data":nil}, sch.Err
	}
	return map[string][]*model.Tier{"data":sch.Data.([]*model.Tier)}, nil

}
