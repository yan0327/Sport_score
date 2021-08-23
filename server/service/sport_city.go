package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateSportCity 创建SportCity记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateSportCity(sportCity model.SportCity) (err error) {
	err = global.GVA_DB.Create(&sportCity).Error
	return err
}

// DeleteSportCity 删除SportCity记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSportCity(sportCity model.SportCity) (err error) {
	err = global.GVA_DB.Delete(&sportCity).Error
	return err
}

// DeleteSportCityByIds 批量删除SportCity记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSportCityByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SportCity{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSportCity 更新SportCity记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateSportCity(sportCity model.SportCity) (err error) {
	err = global.GVA_DB.Save(&sportCity).Error
	return err
}

// GetSportCity 根据id获取SportCity记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSportCity(id uint) (err error, sportCity model.SportCity) {
	err = global.GVA_DB.Where("id = ?", id).First(&sportCity).Error
	return
}

// GetSportCityInfoList 分页获取SportCity记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSportCityInfoList(info request.SportCitySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SportCity{})
    var sportCitys []model.SportCity
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sportCitys).Error
	return err, sportCitys, total
}
