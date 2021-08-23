// 自动生成模板SportSchool
package model

import (
	"gin-vue-admin/global"
)

// SportSchool 结构体
// 如果含有time.Time 请自行import time包
type SportSchool struct {
	global.GVA_MODEL
	Name               string              `json:"name" form:"name" gorm:"column:name;comment:校名;type:varchar(191);"`
	SportCityID        int                 `json:"sportCityID" form:"sportCityID" gorm:"column:sportCityID;comment:市ID;type:bigint;size:20;"`
	SportStudents      []SportStudent      `json:"sportStudent" form:"sportStudent" gorm:"comment:学生"`             // 市名
	SportStudentScores []SportStudentScore `json:"sportStudentScore" form:"sportStudentScore" gorm:"comment:学生成绩"` // 市名
}

// TableName SportSchool 表名
func (SportSchool) TableName() string {
	return "sportSchool"
}
