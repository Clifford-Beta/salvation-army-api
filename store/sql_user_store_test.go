package store

import (
	//"testing"
	"fmt"
	"salvation-army-api/model"
	"testing"
	"time"
)

var pass string = "BT1290%R65snh8"
var user = model.User{
	//Id:1,
	Name:     "Willyss Beta",
	Email:    "betawillys@gmail.com",
	Password: "12345",
	Status:   1,
	DateAdd:  time.Now(),
}

func TestSqlUserStore_Save(t *testing.T) {
	userStore := SqlUserStore{Database}

	//go TimeOut()
	me := <-userStore.Save(&user)
	if me.Err != nil {
		t.Fatal("Saving insurer user failed with", me.Err)
	}
}

func TestSqlUserStore_Get(t *testing.T) {
	userStore := SqlUserStore{Database}

	//go TimeOut()
	me := <-userStore.Get(1)
	if me.Err != nil {
		t.Fatal("Getting  insurer user failed with", me.Err)
	}
	fmt.Println(me.Data)
}

func TestSqlUserStore_GetManyt(t *testing.T) {
	userStore := SqlUserStore{Database}

	//go TimeOut()
	me := <-userStore.GetMany()
	if me.Err != nil {
		t.Fatal("Getting  insurer user failed with", me.Err)
	}
	for _, v := range me.Data.([]*model.User) {
		fmt.Println(*v)

	}
	//fmt.Println(me.Data.([] *model.User))
}

//func TestSqlUserStore_GetByEmail(t *testing.T) {
//	userStore := SqlUserStore{dataStore}
//	go TimeOut()
//	select {
//	case me := <-userStore.GetByEmail(user.Email):
//		if me.Err != nil {
//			t.Fatal("Getting user by email failed with", me.Err)
//		}
//	case <-timeout:
//		t.Fatal("Getting user by email  failed with timeout")
//	}
//
//}
//func TestSqlUserStore_Update(t *testing.T) {
//	userStore := SqlUserStore{dataStore}
//	user.Name = "Test User Updated"
//	go TimeOut()
//	select {
//	case me := <-userStore.Update(&user):
//		//t.Error("Data",me.Data,"Error",me.Err)
//		if me.Err != nil {
//			t.Fatal("Updating user failed with", me.Err)
//		}
//	case <-timeout:
//		t.Fatal("Updating user failed with timeout")
//	}
//}
//func TestSqlUserStore_GetByPhone(t *testing.T) {
//
//	userStore := SqlUserStore{dataStore}
//	go TimeOut()
//	select {
//	case me := <-userStore.GetByPhone(user.Phone):
//		if me.Err != nil {
//			t.Fatal("Getting user by phone failed with", me.Err)
//		}
//	case <-timeout:
//		t.Fatal("Getting user by phone failed with timeout")
//	}
//
//}
//func TestSqlUserStore_Get(t *testing.T) {
//	userStore := SqlUserStore{dataStore}
//	go TimeOut()
//	select {
//	case me := <-userStore.Get(user.Id):
//		t.Error(me.Data.(model.User))
//		if me.Err != nil {
//			t.Fatal("Getting user by id failed with", me.Err)
//		}
//	case <-timeout:
//		t.Fatal("Getting user by id failed with timeout")
//	}
//
//}
////
//func TestSqlUserStore_GetByPhoneAndPassword(t *testing.T) {
//
//	userStore := SqlUserStore{dataStore}
//	go TimeOut()
//	select {
//	case me := <-userStore.GetByPhoneAndPassword(user.Phone, pass):
//		if me.Err != nil {
//			t.Fatal("Getting user by phone and password failed with", me.Err, me.Data)
//		}
//	case <-timeout:
//		t.Fatal("Getting user by phone and password failed with timeout")
//	}
//
//}
//
//func TestSqlUserStore_GetByEmailAndPassword(t *testing.T) {
//
//	userStore := SqlUserStore{dataStore}
//	go TimeOut()
//	select {
//	case me := <-userStore.GetByEmailAndPassword(user.Email, pass):
//		if me.Err != nil {
//			t.Fatal("Getting insurer by email and password failed with", me.Err)
//		}
//	case <-timeout:
//		t.Fatal("Getting insurer by email and password failed with timeout")
//	}
//
//}
//
//func TestSqlUserStore_Delete(t *testing.T) {
//	userStore := SqlUserStore{dataStore}
//	go TimeOut()
//	select {
//	case me := <-userStore.Delete(&user):
//		if me.Err != nil {
//			t.Fatal("Deleting insurer user failed with", me.Err, me.Data)
//		}
//	case <-timeout:
//		t.Fatal("Deleting insurer user failed with timeout")
//	}
//}

//func TestSqlInsurerUserStore_GetMany(t *testing.T) {
//	memberStore := SqlUserStore{dataStore}
//	go TimeOut()
//	select {
//	case me := <-memberStore.GetMany():
//		if me.Err != nil {
//			t.Fatal("Getting user by id failed with",me.Err)
//		}
//	case <-timeout:
//		t.Fatal("Getting user by id failed with timeout")
//	}
//}
