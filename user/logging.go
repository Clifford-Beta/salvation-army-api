package user

import (
	log "github.com/sirupsen/logrus"
	"salvation-army-api/model"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   UserService
}

func (mw LoggingMiddleware) Create(user model.User) (output *model.User, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  user,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "user ", "method = ", "create")

	}(time.Now())
	output, err = mw.Next.Create(user)
	return
}

func (mw LoggingMiddleware) Update(user model.User) (output bool, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  user,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "user ", "method = ", "update")

	}(time.Now())
	output, err = mw.Next.Update(user)
	return
}

func (mw LoggingMiddleware) GetOne(id int) (output model.User, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "user ", "method = ", "getone")

	}(time.Now())
	output, err = mw.Next.GetOne(id)
	return
}

func (mw LoggingMiddleware) Login(email, password string) (output model.User, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  email + password,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "user ", "method = ", "login")

	}(time.Now())
	output, err = mw.Next.Login(email, password)
	return
}

func (mw LoggingMiddleware) GetAll() (output map[string][]*model.User, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  "",
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "user ", "method = ", "getall")

	}(time.Now())
	output, err = mw.Next.GetAll()
	return
}
