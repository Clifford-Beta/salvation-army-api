package user

import (
	"time"
	"github.com/go-kit/kit/log"
	"salv_prj/model"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   UserService
}

func (mw LoggingMiddleware) Create(user model.User) (output *model.User, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "create",
			"input", user.ToJson() ,
			"output", output.ToJson(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.Next.Create(user)
	return
}

func (mw LoggingMiddleware) GetOne(id int) (output model.User, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "getone",
			"input", id ,
			"output", output.ToJson(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.Next.GetOne(id)
	return
}

func (mw LoggingMiddleware) GetAll() (output []*model.User, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "getall",
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.Next.GetAll()
	return
}
