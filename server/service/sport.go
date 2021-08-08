package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateSport 创建Sport记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateSport(sport model.Sport) (err error) {
	err = global.GVA_DB.Create(&sport).Error
	return err
}

// DeleteSport 删除Sport记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSport(sport model.Sport) (err error) {
	err = global.GVA_DB.Delete(&sport).Error
	return err
}

// DeleteSportByIds 批量删除Sport记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSportByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Sport{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSport 更新Sport记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateSport(sport model.Sport) (err error) {
	err = global.GVA_DB.Save(&sport).Error
	return err
}

// UpdateSportByHash 更新Sport记录
// Author [piexlmax](https://github.com/piexlmax)
/**/
func FindSportByHash(Hash256 string) (err error, sport model.Sport) {
	//err = global.GVA_DB.Where("Hash256 = ?", Hash256).First(&sport).Model("TransHash", TransHash)
	err = global.GVA_DB.Model(&sport).Where("Hash256 = ?", Hash256).First(&sport).Error
	return
}

// GetSport 根据id获取Sport记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSport(id uint) (err error, sport model.Sport) {
	err = global.GVA_DB.Where("id = ?", id).First(&sport).Error
	return
}

// GetSportByHash 根据id获取Sport记录
// Author [piexlmax](https://github.com/piexlmax)
/*func GetSportByHash(Hash256 string) (err error, sport model.Sport) {
	err = global.GVA_DB.Where("Hash256 = ?", Hash256).First(&sport).Error
	return
}*/

// GetSportInfoList 分页获取Sport记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSportInfoList(info request.SportSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Sport{})
	var sports []model.Sport
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.School != "" {
		db = db.Where("`school` = ?", info.School)
	}
	if info.Class != "" {
		db = db.Where("`class` = ?", info.Class)
	}
	if info.Testid != "" {
		db = db.Where("`testid` = ?", info.Testid)
	}
	if info.Name != "" {
		db = db.Where("`name` = ?", info.Name)
	}
	if info.Sex != "" {
		db = db.Where("`sex` = ?", info.Sex)
	}
	if info.TotalScore != 0 {
		db = db.Where("`total_score` = ?", info.TotalScore)
	}
	if info.ProcesseValuation != 0 {
		db = db.Where("`processe_valuation` = ?", info.ProcesseValuation)
	}
	if info.Grade != "" {
		db = db.Where("`grade` = ?", info.Grade)
	}
	if info.Itemone != "" {
		db = db.Where("`itemone` = ?", info.Itemone)
	}
	if info.Gradeone != 0 {
		db = db.Where("`gradeone` = ?", info.Gradeone)
	}
	if info.ScoreOne != 0 {
		db = db.Where("`score_one` = ?", info.ScoreOne)
	}
	if info.ItemTwo != "" {
		db = db.Where("`item_two` = ?", info.ItemTwo)
	}
	if info.GradeTwo != 0 {
		db = db.Where("`grade_two` = ?", info.GradeTwo)
	}
	if info.ScoreTwo != 0 {
		db = db.Where("`score_two` = ?", info.ScoreTwo)
	}
	if info.ItemThree != "" {
		db = db.Where("`item_three` = ?", info.ItemThree)
	}
	if info.GradeThree != 0 {
		db = db.Where("`grade_three` = ?", info.GradeThree)
	}
	if info.ScoreThree != 0 {
		db = db.Where("`score_three` = ?", info.ScoreThree)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sports).Error
	return err, sports, total
}
