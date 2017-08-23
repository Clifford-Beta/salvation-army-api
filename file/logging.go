package file

import (
	log "github.com/sirupsen/logrus"
	"salv_prj/model"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   FileService
}

func (mw LoggingMiddleware) Create(file model.File) (output *model.File, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  file,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "file ", "method = ", "create")

	}(time.Now())
	output, err = mw.Next.Create(file)
	return
}
func (mw LoggingMiddleware) CreateType(file model.FileType) (output *model.FileType, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  file,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "file ", "method = ", "create_type")

	}(time.Now())
	output, err = mw.Next.CreateType(file)
	return
}

func (mw LoggingMiddleware) GetOne(id int) (output model.File, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "file ", "method = ", "getone")

	}(time.Now())
	output, err = mw.Next.GetOne(id)
	return
}

func (mw LoggingMiddleware) GetOneType(id int) (output model.FileType, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "file ", "method = ", "get_one_type")

	}(time.Now())
	output, err = mw.Next.GetOneType(id)
	return
}

func (mw LoggingMiddleware) GetAll() (output map[string][]*model.File, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  "",
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "file ", "method = ", "getall")

	}(time.Now())
	output, err = mw.Next.GetAll()
	return
}

func (mw LoggingMiddleware) GetAllTypes() (output map[string][]*model.FileType, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  "",
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "file ", "method = ", "getall")

	}(time.Now())
	output, err = mw.Next.GetAllTypes()
	return
}
