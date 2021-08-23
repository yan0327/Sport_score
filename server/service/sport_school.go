package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateSportSchool 创建SportSchool记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateSportSchool(sportSchool model.SportSchool) (err error) {
	err = global.GVA_DB.Create(&sportSchool).Error
	return err
}

// DeleteSportSchool 删除SportSchool记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSportSchool(sportSchool model.SportSchool) (err error) {
	err = global.GVA_DB.Delete(&sportSchool).Error
	return err
}

// DeleteSportSchoolByIds 批量删除SportSchool记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSportSchoolByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SportSchool{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSportSchool 更新SportSchool记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateSportSchool(sportSchool model.SportSchool) (err error) {
	err = global.GVA_DB.Save(&sportSchool).Error
	return err
}

// GetSportSchool 根据id获取SportSchool记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSportSchool(id uint) (err error, sportSchool model.SportSchool) {
	err = global.GVA_DB.Where("id = ?", id).First(&sportSchool).Error
	return
}

// GetSportSchoolInfoList 分页获取SportSchool记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSportSchoolInfoList(info request.SportSchoolSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SportSchool{})
    var sportSchools []model.SportSchool
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sportSchools).Error
	return err, sportSchools, total
}
