package model

import (
	"gin-vue-admin/global"
)

// Sport 结构体
// 如果含有time.Time 请自行import time包
type SportSchool struct {
	global.GVA_MODEL
	Name               string              `json:"name" form:"name" gorm:"column:name;comment:校名"`
	SportCityID        uint                `json:"sportCityID" form:"sportCityID" gorm:"column:sportCityID;comment:市ID"`
	SportStudents      []SportStudent      `json:"sportStudent" form:"sportStudent" gorm:"comment:学生"`             // 市名
	SportStudentScores []SportStudentScore `json:"sportStudentScore" form:"sportStudentScore" gorm:"comment:学生成绩"` // 市名
}

// TableName Sport 表名
func (SportSchool) TableName() string {
	return "sportSchool"
}
