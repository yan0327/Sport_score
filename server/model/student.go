// 自动生成模板Student
package model

import (
	"gin-vue-admin/global"
)

// Student 结构体
// 如果含有time.Time 请自行import time包
type Student struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名"`
      Age  int `json:"age" form:"age" gorm:"column:age;comment:年龄;type:int;size:3;"`
      Sex  int `json:"sex" form:"sex" gorm:"column:sex;comment:性别;type:int;"`
}


// TableName Student 表名
func (Student) TableName() string {
  return "student"
}

