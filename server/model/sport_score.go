// 自动生成模板Sport_score
package model

import (
	"gin-vue-admin/global"
)

// Sport_score 结构体
// 如果含有time.Time 请自行import time包
type Sport_score struct {
      global.GVA_MODEL
      School  string `json:"school" form:"school" gorm:"column:school;comment:学校;type:varchar;"`
      Class  string `json:"class" form:"class" gorm:"column:class;comment:班级"`
      Testid  string `json:"testid" form:"testid" gorm:"column:testid;comment:考号"`
      Sex  *bool `json:"sex" form:"sex" gorm:"column:sex;comment:性别;type:tinyint"`
      Total_score  float64 `json:"total_score" form:"total_score" gorm:"column:total_score;comment:总分"`
      Process_evaluation  float64 `json:"process_evaluation" form:"process_evaluation" gorm:"column:process_evaluation;comment:过程评价;type:float;"`
      Grade  string `json:"grade" form:"grade" gorm:"column:grade;comment:等级"`
      Item_one  string `json:"item_one" form:"item_one" gorm:"column:item_one;comment:一类项目"`
      Grade_one  float64 `json:"grade_one" form:"grade_one" gorm:"column:grade_one;comment:成绩1;type:float;"`
      Score_one  float64 `json:"score_one" form:"score_one" gorm:"column:score_one;comment:分数1;type:float;"`
      Item_two  string `json:"item_two" form:"item_two" gorm:"column:item_two;comment:二类项目"`
      Grade_two  float64 `json:"grade_two" form:"grade_two" gorm:"column:grade_two;comment:成绩2;type:float;"`
      Score_two  float64 `json:"score_two" form:"score_two" gorm:"column:score_two;comment:分数2;type:float;"`
      Item_three  string `json:"item_three" form:"item_three" gorm:"column:item_three;comment:三类项目"`
      Grade_three  float64 `json:"grade_three" form:"grade_three" gorm:"column:grade_three;comment:成绩3;type:float;"`
      Score_three  float64 `json:"score_three" form:"score_three" gorm:"column:score_three;comment:分数3;type:float;"`
}


// TableName Sport_score 表名
func (Sport_score) TableName() string {
  return "sport_score"
}

