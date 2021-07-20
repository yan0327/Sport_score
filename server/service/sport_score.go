package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateSport_score 创建Sport_score记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateSport_score(sport_score model.Sport_score) (err error) {
	err = global.GVA_DB.Create(&sport_score).Error
	return err
}

// DeleteSport_score 删除Sport_score记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSport_score(sport_score model.Sport_score) (err error) {
	err = global.GVA_DB.Delete(&sport_score).Error
	return err
}

// DeleteSport_scoreByIds 批量删除Sport_score记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSport_scoreByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Sport_score{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSport_score 更新Sport_score记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateSport_score(sport_score model.Sport_score) (err error) {
	err = global.GVA_DB.Save(&sport_score).Error
	return err
}

// GetSport_score 根据id获取Sport_score记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSport_score(id uint) (err error, sport_score model.Sport_score) {
	err = global.GVA_DB.Where("id = ?", id).First(&sport_score).Error
	return
}

// GetSport_scoreInfoList 分页获取Sport_score记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSport_scoreInfoList(info request.Sport_scoreSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Sport_score{})
    var sport_scores []model.Sport_score
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.School != "" {
        db = db.Where("`school` = ?",info.School)
    }
    if info.Class != "" {
        db = db.Where("`class` = ?",info.Class)
    }
    if info.Testid != "" {
        db = db.Where("`testid` = ?",info.Testid)
    }
    if info.Name != "" {
        db = db.Where("`name` = ?",info.Name)
    }
    if info.Sex != nil {
        db = db.Where("`sex` = ?",info.Sex)
    }
    if info.TotalScore != 0 {
        db = db.Where("`total_score` = ?",info.TotalScore)
    }
    if info.ProcesseValuation != 0 {
        db = db.Where("`processe_valuation` = ?",info.ProcesseValuation)
    }
    if info.Grade != "" {
        db = db.Where("`grade` = ?",info.Grade)
    }
    if info.Itemone != "" {
        db = db.Where("`itemone` = ?",info.Itemone)
    }
    if info.Gradeone != 0 {
        db = db.Where("`gradeone` = ?",info.Gradeone)
    }
    if info.ScoreOne != 0 {
        db = db.Where("`score_one` = ?",info.ScoreOne)
    }
    if info.ItemTwo != "" {
        db = db.Where("`item_two` = ?",info.ItemTwo)
    }
    if info.ItemThree != "" {
        db = db.Where("`item_three` = ?",info.ItemThree)
    }
    if info.ScoreThree != 0 {
        db = db.Where("`score_three` = ?",info.ScoreThree)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sport_scores).Error
	return err, sport_scores, total
}
