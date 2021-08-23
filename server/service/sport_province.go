package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateSportProvince 创建SportProvince记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateSportProvince(sportProvince model.SportProvince) (err error) {
	err = global.GVA_DB.Create(&sportProvince).Error
	return err
}

// DeleteSportProvince 删除SportProvince记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSportProvince(sportProvince model.SportProvince) (err error) {
	err = global.GVA_DB.Delete(&sportProvince).Error
	return err
}

// DeleteSportProvinceByIds 批量删除SportProvince记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSportProvinceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SportProvince{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSportProvince 更新SportProvince记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateSportProvince(sportProvince model.SportProvince) (err error) {
	err = global.GVA_DB.Save(&sportProvince).Error
	return err
}

// GetSportProvince 根据id获取SportProvince记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSportProvince(id uint) (err error, sportProvince model.SportProvince) {
	err = global.GVA_DB.Where("id = ?", id).First(&sportProvince).Error
	return
}

// GetSportProvinceInfoList 分页获取SportProvince记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSportProvinceInfoList(info request.SportProvinceSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SportProvince{})
    var sportProvinces []model.SportProvince
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sportProvinces).Error
	return err, sportProvinces, total
}
