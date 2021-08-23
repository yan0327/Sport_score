// 自动生成模板SportStudentScore
package model

import (
	"gin-vue-admin/global"
)

// SportStudentScore 结构体
// 如果含有time.Time 请自行import time包
type SportStudentScore struct {
	global.GVA_MODEL
	Class          string `json:"class" form:"class" gorm:"column:class;comment:大类;type:varchar(191);"`
	Name           string `json:"name" form:"name" gorm:"column:name;comment:方法名称;type:varchar(191);"`
	Score          string `json:"score" form:"score" gorm:"column:score;comment:分数;type:varchar(191);"`
	SportSchoolID  int    `json:"sportSchoolID" form:"sportSchoolID" gorm:"column:sportSchoolID;comment:学校ID;type:bigint;size:20;"`
	SportStudentID int    `json:"sportStudentID" form:"sportStudentID" gorm:"column:sportStudentID;comment:学生ID;type:bigint;size:20;"`
	Testdata       string `json:"testdata" form:"testdata" gorm:"column:testdata;comment:测试数据;type:varchar(191);"`
}

// TableName SportStudentScore 表名
func (SportStudentScore) TableName() string {
	return "sportStudentScore"
}
