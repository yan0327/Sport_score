// 自动生成模板Sport_score
package model

import (
	"gin-vue-admin/global"
)

// Sport_score 结构体
// 如果含有time.Time 请自行import time包
type Sport_score struct {
      global.GVA_MODEL
      School  string `json:"school" form:"school" gorm:"column:school;comment:学校"`
      Class  string `json:"class" form:"class" gorm:"column:class;comment:班级"`
      Testid  string `json:"testid" form:"testid" gorm:"column:testid;comment:考号"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名"`
      Sex  *bool `json:"sex" form:"sex" gorm:"column:sex;comment:性别;type:tinyint"`
      TotalScore  float64 `json:"totalScore" form:"totalScore" gorm:"column:total_score;comment:总分"`
      ProcesseValuation  float64 `json:"processeValuation" form:"processeValuation" gorm:"column:processe_valuation;comment:过程评价"`
      Grade  string `json:"grade" form:"grade" gorm:"column:grade;comment:等级"`
      Itemone  string `json:"itemone" form:"itemone" gorm:"column:itemone;comment:一类项目"`
      Gradeone  float64 `json:"gradeone" form:"gradeone" gorm:"column:gradeone;comment:成绩1"`
      ScoreOne  float64 `json:"scoreOne" form:"scoreOne" gorm:"column:score_one;comment:分数1;type:float;"`
      ItemTwo  string `json:"itemTwo" form:"itemTwo" gorm:"column:item_two;comment:二类项目"`
      GradeTwo  float64 `json:"gradeTwo" form:"gradeTwo" gorm:"column:grade_two;comment:成绩2"`
      ScoreTwo  float64 `json:"scoreTwo" form:"scoreTwo" gorm:"column:score_two;comment:分数2;type:float;"`
      ItemThree  string `json:"itemThree" form:"itemThree" gorm:"column:item_three;comment:三类项目"`
      GradeThree  float64 `json:"gradeThree" form:"gradeThree" gorm:"column:grade_three;comment:成绩3;type:float;"`
      ScoreThree  float64 `json:"scoreThree" form:"scoreThree" gorm:"column:score_three;comment:分数3;type:float;"`
}


// TableName Sport_score 表名
func (Sport_score) TableName() string {
  return "sport_score"
}

