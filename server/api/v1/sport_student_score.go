package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateSportStudentScore 创建SportStudentScore
// @Tags SportStudentScore
// @Summary 创建SportStudentScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudentScore true "创建SportStudentScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportStudentScore/createSportStudentScore [post]
func CreateSportStudentScore(c *gin.Context) {
	var sportStudentScore model.SportStudentScore
	_ = c.ShouldBindJSON(&sportStudentScore)
	if err := service.CreateSportStudentScore(sportStudentScore); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSportStudentScore 删除SportStudentScore
// @Tags SportStudentScore
// @Summary 删除SportStudentScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudentScore true "删除SportStudentScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportStudentScore/deleteSportStudentScore [delete]
func DeleteSportStudentScore(c *gin.Context) {
	var sportStudentScore model.SportStudentScore
	_ = c.ShouldBindJSON(&sportStudentScore)
	if err := service.DeleteSportStudentScore(sportStudentScore); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSportStudentScoreByIds 批量删除SportStudentScore
// @Tags SportStudentScore
// @Summary 批量删除SportStudentScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SportStudentScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sportStudentScore/deleteSportStudentScoreByIds [delete]
func DeleteSportStudentScoreByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSportStudentScoreByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSportStudentScore 更新SportStudentScore
// @Tags SportStudentScore
// @Summary 更新SportStudentScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudentScore true "更新SportStudentScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sportStudentScore/updateSportStudentScore [put]
func UpdateSportStudentScore(c *gin.Context) {
	var sportStudentScore model.SportStudentScore
	_ = c.ShouldBindJSON(&sportStudentScore)
	if err := service.UpdateSportStudentScore(sportStudentScore); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSportStudentScore 用id查询SportStudentScore
// @Tags SportStudentScore
// @Summary 用id查询SportStudentScore
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudentScore true "用id查询SportStudentScore"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sportStudentScore/findSportStudentScore [get]
func FindSportStudentScore(c *gin.Context) {
	var sportStudentScore model.SportStudentScore
	_ = c.ShouldBindQuery(&sportStudentScore)
	if err, resportStudentScore := service.GetSportStudentScore(sportStudentScore.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resportStudentScore": resportStudentScore}, c)
	}
}

// GetSportStudentScoreList 分页获取SportStudentScore列表
// @Tags SportStudentScore
// @Summary 分页获取SportStudentScore列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SportStudentScoreSearch true "分页获取SportStudentScore列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportStudentScore/getSportStudentScoreList [get]
func GetSportStudentScoreList(c *gin.Context) {
	var pageInfo request.SportStudentScoreSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSportStudentScoreInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
