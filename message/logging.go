package message

import (
	log "github.com/sirupsen/logrus"
	"salvation-army-api/model"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   MessageService
}

func (mw LoggingMiddleware) Create(message model.Message) (output *model.Message, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  message,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "message ", "method = ", "create")

	}(time.Now())
	output, err = mw.Next.Create(message)
	return
}

func (mw LoggingMiddleware) Update(message model.Message) (output bool, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  message,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "message ", "method = ", "update")

	}(time.Now())
	output, err = mw.Next.Update(message)
	return
}

func (mw LoggingMiddleware) Delete(message model.Message) (output bool, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  message,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "message ", "method = ", "delete")

	}(time.Now())
	output, err = mw.Next.Delete(message)
	return
}

func (mw LoggingMiddleware) GetOne(id int) (output model.Message, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "message ", "method = ", "getone")

	}(time.Now())
	output, err = mw.Next.GetOne(id)
	return
}

func (mw LoggingMiddleware) GetAll(user int) (output map[string][]model.MessageResult, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  "",
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "message ", "method = ", "getall")

	}(time.Now())
	output, err = mw.Next.GetAll(user)
	return
}
