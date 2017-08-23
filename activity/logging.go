package activity

import (
	log "github.com/sirupsen/logrus"
	"salvation-army-api/model"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   ActivityService
}

func (mw LoggingMiddleware) Create(activity model.ExtraCurricular) (output *model.ExtraCurricular, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  activity,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "activity ", "method = ", "create")
	}(time.Now())
	output, err = mw.Next.Create(activity)
	return
}

func (mw LoggingMiddleware) CreateLevel(level model.ExtraCurricularLevel) (output *model.ExtraCurricularLevel, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  level,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "activity ", "method = ", "create_level")
	}(time.Now())
	output, err = mw.Next.CreateLevel(level)
	return
}

func (mw LoggingMiddleware) RecordPerformance(performance model.ExtraCurricularActivity) (output *model.ExtraCurricularActivity, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  performance,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "activity ", "method = ", "record_performance")
	}(time.Now())
	output, err = mw.Next.RecordPerformance(performance)
	return
}

func (mw LoggingMiddleware) GetOneActivity(id int) (output model.ExtraCurricular, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "activity ", "method = ", "get_one")
	}(time.Now())
	output, err = mw.Next.GetOneActivity(id)
	return
}

func (mw LoggingMiddleware) GetOneLevel(id int) (output model.ExtraCurricularLevel, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "activity ", "method = ", "get_one_level")
	}(time.Now())
	output, err = mw.Next.GetOneLevel(id)
	return
}

func (mw LoggingMiddleware) GetOnePerformance(id int) (output model.ExtraCurricularActivity, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "activity ", "method = ", "get_one_performance")
	}(time.Now())
	output, err = mw.Next.GetOnePerformance(id)
	return
}

func (mw LoggingMiddleware) GetAllActivities() (output map[string][]*model.ExtraCurricular, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "activity ", "method = ", "get_all_activities")
	}(time.Now())
	output, err = mw.Next.GetAllActivities()
	return
}

func (mw LoggingMiddleware) GetAllLevels() (output map[string][]*model.ExtraCurricularLevel, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "activity ", "method = ", "get_all_levels")
	}(time.Now())
	output, err = mw.Next.GetAllLevels()
	return
}

func (mw LoggingMiddleware) GetAllPerformances() (output map[string][]*model.ExtraCurricularActivity, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "activity ", "method = ", "get_all_performances")
	}(time.Now())
	output, err = mw.Next.GetAllPerformances()
	return
}
