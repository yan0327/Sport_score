package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateSportStudent 创建SportStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateSportStudent(sportStudent model.SportStudent) (err error) {
	err = global.GVA_DB.Create(&sportStudent).Error
	return err
}

// DeleteSportStudent 删除SportStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSportStudent(sportStudent model.SportStudent) (err error) {
	err = global.GVA_DB.Delete(&sportStudent).Error
	return err
}

// DeleteSportStudentByIds 批量删除SportStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSportStudentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SportStudent{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSportStudent 更新SportStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateSportStudent(sportStudent model.SportStudent) (err error) {
	err = global.GVA_DB.Save(&sportStudent).Error
	return err
}

// GetSportStudent 根据id获取SportStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSportStudent(id uint) (err error, sportStudent model.SportStudent) {
	err = global.GVA_DB.Where("id = ?", id).First(&sportStudent).Error
	return
}

// GetSportStudentInfoList 分页获取SportStudent记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSportStudentInfoList(info request.SportStudentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SportStudent{})
    var sportStudents []model.SportStudent
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sportStudents).Error
	return err, sportStudents, total
}
