package category

import (
	log "github.com/sirupsen/logrus"
	"salvation-army-api/model"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   CategoryService
}

func (mw LoggingMiddleware) Create(category model.Category) (output *model.Category, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  category,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "category ", "method = ", "create")
	}(time.Now())
	output, err = mw.Next.Create(category)
	return
}

func (mw LoggingMiddleware) CreateTier(category model.Tier) (output *model.Tier, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  category,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "category ", "method = ", "create_tier")
	}(time.Now())
	output, err = mw.Next.CreateTier(category)
	return
}

func (mw LoggingMiddleware) GetOne(id int) (output model.Category, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "category ", "method = ", "get_one")
	}(time.Now())
	output, err = mw.Next.GetOne(id)
	return
}

func (mw LoggingMiddleware) GetOneTier(id int) (output model.Tier, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "category ", "method = ", "get_one_tier")
	}(time.Now())
	output, err = mw.Next.GetOneTier(id)
	return
}

func (mw LoggingMiddleware) GetAll() (output []*model.Category, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "category ", "method = ", "get_all")
	}(time.Now())
	output, err = mw.Next.GetAll()
	return
}

func (mw LoggingMiddleware) GetAllTiers() (output []*model.Tier, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "category ", "method = ", "get_all_tiers")
	}(time.Now())
	output, err = mw.Next.GetAllTiers()
	return
}
