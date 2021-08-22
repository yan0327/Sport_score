package model

import (
	"gin-vue-admin/global"
)

// Sport 结构体
// 如果含有time.Time 请自行import time包
type SportProvince struct {
	global.GVA_MODEL
	Name       string      `json:"name" form:"name" gorm:"column:name;comment:省名"`
	SportCitys []SportCity `json:"sportCity" form:"sportCity" gorm:"comment:市名"` // 市名
}

// TableName Sport 表名
func (SportProvince) TableName() string {
	return "sportProvince"
}
