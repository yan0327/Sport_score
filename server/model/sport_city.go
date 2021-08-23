// 自动生成模板SportCity
package model

import (
	"gin-vue-admin/global"
)

// SportCity 结构体
// 如果含有time.Time 请自行import time包
type SportCity struct {
	global.GVA_MODEL
	Name            string        `json:"name" form:"name" gorm:"column:name;comment:市名;type:varchar(191);"`
	SportProvinceID int           `json:"sportProvinceID" form:"sportProvinceID" gorm:"column:sportProvinceID;comment:省ID;type:bigint;size:20;"`
	SportSchools    []SportSchool `json:"sportSchool" form:"sportSchool" gorm:"comment:校名"` // 市名
}

// TableName SportCity 表名
func (SportCity) TableName() string {
	return "sportCity"
}
