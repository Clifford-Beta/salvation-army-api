package school

import (
	"time"
	//"github.com/go-kit/kit/log"
	log "github.com/sirupsen/logrus"
	"salvation-army-api/model"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   SchoolService
}

func (mw LoggingMiddleware) Create(school model.School) (output *model.School, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  school,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "school ", "method = ", "create")
	}(time.Now())
	output, err = mw.Next.Create(school)
	return
}

func (mw LoggingMiddleware) Update(school model.School) (output UpdateResponse, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  school,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "school ", "method = ", "update")
	}(time.Now())
	output, err = mw.Next.Update(school)
	return
}

func (mw LoggingMiddleware) Delete(school model.School) (output bool, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  school,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "school ", "method = ", "delete")
	}(time.Now())
	output, err = mw.Next.Delete(school)
	return
}

func (mw LoggingMiddleware) GetOne(id int) (output model.School, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "school ", "method = ", "getone")
	}(time.Now())
	output, err = mw.Next.GetOne(id)
	return
}

func (mw LoggingMiddleware) GetAll() (output map[string][]*model.SchoolResult, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "school ", "method = ", "getall")
	}(time.Now())
	output, err = mw.Next.GetAll()
	return
}

func (mw LoggingMiddleware) RecordPerformance(performance *model.SchoolPerformance) (output *model.SchoolPerformance, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"params": map[string]interface{}{"performance": performance},
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "school ", "method = ", "recordperformance")
	}(time.Now())
	output, err = mw.Next.RecordPerformance(performance)
	return
}

func (mw LoggingMiddleware) GetBestSchool(from, to int) (output model.SchoolPerformanceResult, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"params": map[string]interface{}{"from": from, "to": to},
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "school ", "method = ", "getbestschool")
	}(time.Now())
	output, err = mw.Next.GetBestSchool(from, to)
	return
}

func (mw LoggingMiddleware) RankAllSchools(from, to int) (output map[string][]model.SchoolPerformanceResult, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"params": map[string]interface{}{"from": from, "to": to},
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "school ", "method = ", "rankallschools")
	}(time.Now())
	output, err = mw.Next.RankAllSchools(from, to)
	return
}
