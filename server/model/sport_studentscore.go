// 自动生成模板Sport_score
package model

import "gin-vue-admin/global"

// Sport_score 结构体
// 如果含有time.Time 请自行import time包
type SportStudentScore struct {
	global.GVA_MODEL
	Class          string `json:"class" form:"class" gorm:"column:class;comment:大类"`
	Name           string `json:"name" form:"name" gorm:"column:name;comment:方法名称"`
	Testdata       string `json:"testdata" form:"testdata" gorm:"column:testdata;comment:测试数据"`
	Score          string `json:"score" form:"score" gorm:"column:score;comment:分数"`
	SportStudentID uint   `json:"sportStudentID" form:"sportStudentID" gorm:"column:sportStudentID;comment:学生ID"`
	SportSchoolID  uint   `json:"sportSchoolID" form:"sportSchoolID" gorm:"column:sportSchoolID;comment:学校ID"`
}

// TableName Sport_score 表名
func (SportStudentScore) TableName() string {
	return "sportStudentScore"
}
