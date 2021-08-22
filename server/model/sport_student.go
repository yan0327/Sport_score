// 自动生成模板Sport_score
package model

import "gin-vue-admin/global"

// Sport_score 结构体
// 如果含有time.Time 请自行import time包
type SportStudent struct {
	global.GVA_MODEL
	Name               string              `json:"name" form:"name" gorm:"column:name;comment:姓名"`
	Sex                string              `json:"sex" form:"sex" gorm:"column:sex;comment:性别;type:varchar(20);"`
	Testid             string              `json:"testid" form:"testid" gorm:"column:testid;comment:学号"`
	Idcard             string              `json:"idcard" form:"idcard" gorm:"column:idcard;comment:身份证"`
	Ethnicity          string              `json:"ethnicity" form:"ethnicity" gorm:"column:ethnicity;comment:民族"`
	PoliticalStatus    string              `json:"politicalstatus" form:"politicalstatus" gorm:"column:politicalstatus;comment:政治面貌"`
	School             string              `json:"school" form:"school" gorm:"column:school;comment:学校"`
	Class              string              `json:"class" form:"class" gorm:"column:class;comment:班级"`
	PhoneNum1          string              `json:"phonenum1" form:"phonenum1" gorm:"column:phonenum1;comment:联系方式"`
	PhoneNum2          string              `json:"phonenum2" form:"phonenum2" gorm:"column:phonenum2;comment:监护人联系方式"`
	SportSchoolID      uint                `json:"sportSchoolID" form:"sportSchoolID" gorm:"column:sportSchoolID;comment:学校ID"`
	SportStudentScores []SportStudentScore `json:"sportStudentScore" form:"sportStudentScore" gorm:"comment:学生成绩"` // 市名
}

// TableName Sport_score 表名
func (SportStudent) TableName() string {
	return "sportStudent"
}
