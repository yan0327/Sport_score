package request

import "gin-vue-admin/model"

type SportStudentScoreSearch struct {
	model.SportStudentScore
	PageInfo
}
