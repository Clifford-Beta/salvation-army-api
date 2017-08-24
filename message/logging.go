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

func (mw LoggingMiddleware) GetAll() (output map[string][]model.Message, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  "",
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "message ", "method = ", "getall")

	}(time.Now())
	output, err = mw.Next.GetAll()
	return
}
