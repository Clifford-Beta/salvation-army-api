package infrastructure

import (
	log "github.com/sirupsen/logrus"
	"salvation-army-api/model"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   InfrastructureService
}

func (mw LoggingMiddleware) Create(inf model.Infrastructure) (output *model.Infrastructure, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  inf,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "infrastructure ", "method = ", "create")

	}(time.Now())
	output, err = mw.Next.Create(inf)
	return
}
func (mw LoggingMiddleware) CreateType(inf model.InfrastructureType) (output *model.InfrastructureType, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  inf,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "infrastructure ", "method = ", "create_type")

	}(time.Now())
	output, err = mw.Next.CreateType(inf)
	return
}

func (mw LoggingMiddleware) GetOne(id int) (output model.InfrastructureResult, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "infrastructure ", "method = ", "getone")

	}(time.Now())
	output, err = mw.Next.GetOne(id)
	return
}

func (mw LoggingMiddleware) GetOneType(id int) (output model.InfrastructureType, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "infrastructure ", "method = ", "get_one_type")

	}(time.Now())
	output, err = mw.Next.GetOneType(id)
	return
}

func (mw LoggingMiddleware) GetAll() (output map[string][]model.InfrastructureResult, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  "",
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "infrastructure ", "method = ", "getall")

	}(time.Now())
	output, err = mw.Next.GetAll()
	return
}

func (mw LoggingMiddleware) GetAllTypes() (output map[string][]model.InfrastructureType, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  "",
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "infrastructure ", "method = ", "getall")

	}(time.Now())
	output, err = mw.Next.GetAllTypes()
	return
}
