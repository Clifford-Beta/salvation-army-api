package user

import (
	"salvation-army-api/model"
	"salvation-army-api/store"
	"time"
)

// StringService provides operations on strings.
type UserService interface {
	Create(model.User) (*model.User, error)
	GetOne(int) (model.User, error)
	GetAll() (map[string][]*model.User, error)
	Login(email, password string) (model.User, error)
	Update(user model.User) (bool, error)
	Delete(user model.User) (bool, error)
}

type Userservice struct{}

func (Userservice) Create(user model.User) (*model.User, error) {
	userStore := store.SqlUserStore{store.Database}
	user.Status = 1
	//user.Password = model.HashPassword(user.Password)
	user.Presave()
	user.DateAdd = time.Now()
	if err := user.Validate(); err != nil {
		return &model.User{}, err
	}
	me := <-userStore.Save(&user)
	if me.Err != nil {
		return &model.User{}, me.Err
	}

	return me.Data.(*model.User).Sanitize(), nil
}

func (Userservice) GetOne(id int) (model.User, error) {
	userStore := store.SqlUserStore{store.Database}
	me := <-userStore.Get(id)
	if me.Err != nil {
		return model.User{}, me.Err
	}
	us := me.Data.(model.User)
	us.Sanitize()
	return us, nil
}

func (Userservice) Update(user model.User) (bool, error) {
	userStore := store.SqlUserStore{store.Database}
	me := <-userStore.Update(&user)
	if me.Err != nil {
		return me.Data.(bool), me.Err
	}
	return me.Data.(bool), nil
}

func (Userservice) Delete(user model.User) (bool, error) {
	userStore := store.SqlUserStore{store.Database}
	me := <-userStore.Delete(&user)
	if me.Err != nil {
		return false, me.Err
	}
	return true, nil
}

func (Userservice) Login(email, password string) (model.User, error) {
	userStore := store.SqlUserStore{store.Database}
	me := <-userStore.GetByEmailAndPassword(email, password)
	if me.Err != nil {
		return model.User{}, me.Err
	}
	return me.Data.(model.User), nil
}
func (Userservice) GetAll() (map[string][]*model.User, error) {
	userStore := store.SqlUserStore{store.Database}
	me := <-userStore.GetMany()
	if me.Err != nil {
		return map[string][]*model.User{"data": []*model.User{}}, me.Err
	}
	return map[string][]*model.User{"data": me.Data.([]*model.User)}, nil
}

//func (Userservice) Update(user model.User) (model.User, error) {
//	return model.User{},nil
//}
