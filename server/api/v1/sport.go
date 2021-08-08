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

// CreateSport 创建Sport
// @Tags Sport
// @Summary 创建Sport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport true "创建Sport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sport/createSport [post]
func CreateSport(c *gin.Context) {
	var sport model.Sport
	_ = c.ShouldBindJSON(&sport)
	if err := service.CreateSport(sport); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSport 删除Sport
// @Tags Sport
// @Summary 删除Sport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport true "删除Sport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sport/deleteSport [delete]
func DeleteSport(c *gin.Context) {
	var sport model.Sport
	_ = c.ShouldBindJSON(&sport)
	if err := service.DeleteSport(sport); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSportByIds 批量删除Sport
// @Tags Sport
// @Summary 批量删除Sport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Sport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sport/deleteSportByIds [delete]
func DeleteSportByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSportByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSport 更新Sport
// @Tags Sport
// @Summary 更新Sport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport true "更新Sport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sport/updateSport [put]
func UpdateSport(c *gin.Context) {
	var sport model.Sport
	_ = c.ShouldBindJSON(&sport)
	if err := service.UpdateSport(sport); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSportByHash 查找Sport
// @Tags Sport
// @Summary 更新Sport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport true "更新Sport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sport/updateSport [put]

func FindSportByHash(c *gin.Context) {
	var sport model.Sport
	_ = c.ShouldBindJSON(&sport)
	if err, resport := service.FindSportByHash(sport.Hash256); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resport": resport}, c)
	}
}

// FindSport 用id查询Sport
// @Tags Sport
// @Summary 用id查询Sport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport true "用id查询Sport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sport/findSport [get]
func FindSport(c *gin.Context) {
	var sport model.Sport
	_ = c.ShouldBindQuery(&sport)
	if err, resport := service.GetSport(sport.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resport": resport}, c)
	}
}



// GetSportList 分页获取Sport列表
// @Tags Sport
// @Summary 分页获取Sport列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SportSearch true "分页获取Sport列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sport/getSportList [get]
func GetSportList(c *gin.Context) {
	var pageInfo request.SportSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSportInfoList(pageInfo); err != nil {
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
