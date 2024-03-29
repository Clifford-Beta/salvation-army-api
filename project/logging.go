package project

import (
	log "github.com/sirupsen/logrus"
	"salvation-army-api/model"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   ProjectService
}

func (mw LoggingMiddleware) Create(project model.Project) (output *model.Project, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  project,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "project ", "method = ", "create")

	}(time.Now())
	output, err = mw.Next.Create(project)
	return
}

func (mw LoggingMiddleware) Update(project model.Project) (output bool, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  project,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "project ", "method = ", "update")

	}(time.Now())
	output, err = mw.Next.Update(project)
	return
}

func (mw LoggingMiddleware) Delete(project model.Project) (output bool, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  project,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "project ", "method = ", "delete")

	}(time.Now())
	output, err = mw.Next.Delete(project)
	return
}

func (mw LoggingMiddleware) GetOne(id int) (output model.Project, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "project ", "method = ", "getone")

	}(time.Now())
	output, err = mw.Next.GetOne(id)
	return
}

func (mw LoggingMiddleware) GetAll() (output map[string][]model.ProjectResult, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  "",
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "project ", "method = ", "getall")

	}(time.Now())
	output, err = mw.Next.GetAll()
	return
}
