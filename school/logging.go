package school

import (
"time"
//"github.com/go-kit/kit/log"
"salv_prj/model"
	log "github.com/sirupsen/logrus"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   SchoolService
}

func (mw LoggingMiddleware) Create(school model.School) (output *model.School, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input": school,
			"output": output,
			"err": err,
			"took": time.Since(begin)}).Info("method = ", "create")
	}(time.Now())
	output, err = mw.Next.Create(school)
	return
}

func (mw LoggingMiddleware) GetOne(id int) (output model.School, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input": id ,
			"output": output.ToJson(),
			"err": err,
			"took": time.Since(begin)}).Info("method = ", "getone",)
	}(time.Now())
	output, err = mw.Next.GetOne(id)
	return
}

func (mw LoggingMiddleware) GetAll() (output []*model.School, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"output": output,
			"err": err,
			"took": time.Since(begin)}).Info("method = ", "getall")
	}(time.Now())
	output, err = mw.Next.GetAll()
	return
}

func (mw LoggingMiddleware) RecordPerformance(performance *model.SchoolPerformance)( output *model.SchoolPerformance, err error)  {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"params":map[string]interface{}{"performance":performance},
			"output": output,
			"err": err,
			"took": time.Since(begin)}).Info("method = ", "recordperformance")
	}(time.Now())
	output, err = mw.Next.RecordPerformance(performance)
	return
}

func (mw LoggingMiddleware) GetBestSchool(from,to int)(output model.SchoolPerformanceResult,err error)  {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"params":map[string]interface{}{"from":from,"to":to},
			"output": output,
			"err": err,
			"took": time.Since(begin)}).Info("method = ","getbestschool")
	}(time.Now())
	output, err = mw.Next.GetBestSchool(from,to)
	return
}

func (mw LoggingMiddleware) RankAllSchools(from,to int)(output []*model.SchoolPerformanceResult,err error)  {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"params":map[string]interface{}{"from":from,"to":to},
			"output": output,
			"err":err,
			"took": time.Since(begin)}).Info("method = ","rankallschools")
	}(time.Now())
	output, err = mw.Next.RankAllSchools(from,to)
	return
}