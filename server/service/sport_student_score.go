package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateSportStudentScore 创建SportStudentScore记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateSportStudentScore(sportStudentScore model.SportStudentScore) (err error) {
	err = global.GVA_DB.Create(&sportStudentScore).Error
	return err
}

// DeleteSportStudentScore 删除SportStudentScore记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSportStudentScore(sportStudentScore model.SportStudentScore) (err error) {
	err = global.GVA_DB.Delete(&sportStudentScore).Error
	return err
}

// DeleteSportStudentScoreByIds 批量删除SportStudentScore记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSportStudentScoreByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SportStudentScore{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSportStudentScore 更新SportStudentScore记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateSportStudentScore(sportStudentScore model.SportStudentScore) (err error) {
	err = global.GVA_DB.Save(&sportStudentScore).Error
	return err
}

// GetSportStudentScore 根据id获取SportStudentScore记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSportStudentScore(id uint) (err error, sportStudentScore model.SportStudentScore) {
	err = global.GVA_DB.Where("id = ?", id).First(&sportStudentScore).Error
	return
}

// GetSportStudentScoreInfoList 分页获取SportStudentScore记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSportStudentScoreInfoList(info request.SportStudentScoreSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SportStudentScore{})
    var sportStudentScores []model.SportStudentScore
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sportStudentScores).Error
	return err, sportStudentScores, total
}
