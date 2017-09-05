package staff

import (
	log "github.com/sirupsen/logrus"
	"salvation-army-api/model"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   StaffService
}

func (mw LoggingMiddleware) AddStaff(staff model.Staff) (output *model.Staff, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  staff,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "add_staff")

	}(time.Now())
	output, err = mw.Next.AddStaff(staff)
	return
}

func (mw LoggingMiddleware) UpdateStaff(staff model.Staff) (output bool, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  staff,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "update_staff")

	}(time.Now())
	output, err = mw.Next.UpdateStaff(staff)
	return
}

func (mw LoggingMiddleware) AddStaffRole(role model.StaffRole) (output *model.StaffRole, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  role,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "add_staff")

	}(time.Now())
	output, err = mw.Next.AddStaffRole(role)
	return
}

func (mw LoggingMiddleware) RetrieveStaff(id int) (output model.Staff, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "retrieve_staff")

	}(time.Now())
	output, err = mw.Next.RetrieveStaff(id)
	return
}

func (mw LoggingMiddleware) RetrieveStaffRole(id int) (output model.StaffRole, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "retrieve_staff_role")

	}(time.Now())
	output, err = mw.Next.RetrieveStaffRole(id)
	return
}

func (mw LoggingMiddleware) GetTeacher(id int) (output model.BestTeacher, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "retrieve_staff_role")

	}(time.Now())
	output, err = mw.Next.GetTeacher(id)
	return
}

func (mw LoggingMiddleware) GetStudent(id int) (output model.BestStudent, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  id,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "retrieve_staff_role")

	}(time.Now())
	output, err = mw.Next.GetStudent(id)
	return
}

func (mw LoggingMiddleware) RetrieveAllStaff() (output map[string][]model.StaffResult, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  "",
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "retrieve_all_staff")

	}(time.Now())
	output, err = mw.Next.RetrieveAllStaff()
	return
}

func (mw LoggingMiddleware) RetrieveAllRoles() (output map[string][]*model.StaffRole, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  "",
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "retrieve_all_staff_roles")

	}(time.Now())
	output, err = mw.Next.RetrieveAllRoles()
	return
}

func (mw LoggingMiddleware) RecordBestPerformingStaff(teacher model.BestTeacher) (output *model.BestTeacher, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  teacher,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "reacord_best_performing_teacher")

	}(time.Now())
	output, err = mw.Next.RecordBestPerformingStaff(teacher)
	return
}

func (mw LoggingMiddleware) UpdateBestPerformingStaff(teacher model.BestTeacher) (output bool, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  teacher,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "update_best_performing_teacher")

	}(time.Now())
	output, err = mw.Next.UpdateBestPerformingStaff(teacher)
	return
}

func (mw LoggingMiddleware) RecordBestPerformingStudent(student model.BestStudent) (output *model.BestStudent, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  student,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "reacord_best_performing_student")

	}(time.Now())
	output, err = mw.Next.RecordBestPerformingStudent(student)
	return
}

func (mw LoggingMiddleware) UpdateBestPerformingStudent(student model.BestStudent) (output bool, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  student,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "update_best_performing_student")

	}(time.Now())
	output, err = mw.Next.UpdateBestPerformingStudent(student)
	return
}

func (mw LoggingMiddleware) RetrieveBestPerformingStudent(from, to int) (output model.BestStudentResult, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  from + to,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "retrieve_best_performing_student")

	}(time.Now())
	output, err = mw.Next.RetrieveBestPerformingStudent(from, to)
	return
}

func (mw LoggingMiddleware) RetrieveBestPerformingStaff(from, to int) (output model.BestTeacherResult, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  from + to,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "retrieve_best_performing_teacher")

	}(time.Now())
	output, err = mw.Next.RetrieveBestPerformingStaff(from, to)
	return
}

func (mw LoggingMiddleware) RankStaffPerformance(from, to int) (output map[string][]*model.BestTeacherResult, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  from + to,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "rank_staff_performance")

	}(time.Now())
	output, err = mw.Next.RankStaffPerformance(from, to)
	return
}

func (mw LoggingMiddleware) RankStudentPerformance(from, to int) (output map[string][]*model.BestStudentResult, err error) {
	defer func(begin time.Time) {
		mw.Logger.WithFields(log.Fields{
			"input":  from + to,
			"output": output,
			"err":    err,
			"took":   time.Since(begin)}).Info("service = ", "staff ", "method = ", "rank_student_performance")

	}(time.Now())
	output, err = mw.Next.RankStudentPerformance(from, to)
	return
}
