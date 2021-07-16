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

// CreateSport_score 创建Sport_score
// @Tags Sport_score
// @Summary 创建Sport_score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport_score true "创建Sport_score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sport_score/createSport_score [post]
func CreateSport_score(c *gin.Context) {
	var sport_score model.Sport_score
	_ = c.ShouldBindJSON(&sport_score)
	if err := service.CreateSport_score(sport_score); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSport_score 删除Sport_score
// @Tags Sport_score
// @Summary 删除Sport_score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport_score true "删除Sport_score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sport_score/deleteSport_score [delete]
func DeleteSport_score(c *gin.Context) {
	var sport_score model.Sport_score
	_ = c.ShouldBindJSON(&sport_score)
	if err := service.DeleteSport_score(sport_score); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSport_scoreByIds 批量删除Sport_score
// @Tags Sport_score
// @Summary 批量删除Sport_score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Sport_score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sport_score/deleteSport_scoreByIds [delete]
func DeleteSport_scoreByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSport_scoreByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSport_score 更新Sport_score
// @Tags Sport_score
// @Summary 更新Sport_score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport_score true "更新Sport_score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sport_score/updateSport_score [put]
func UpdateSport_score(c *gin.Context) {
	var sport_score model.Sport_score
	_ = c.ShouldBindJSON(&sport_score)
	if err := service.UpdateSport_score(sport_score); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSport_score 用id查询Sport_score
// @Tags Sport_score
// @Summary 用id查询Sport_score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Sport_score true "用id查询Sport_score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sport_score/findSport_score [get]
func FindSport_score(c *gin.Context) {
	var sport_score model.Sport_score
	_ = c.ShouldBindQuery(&sport_score)
	if err, resport_score := service.GetSport_score(sport_score.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resport_score": resport_score}, c)
	}
}

// GetSport_scoreList 分页获取Sport_score列表
// @Tags Sport_score
// @Summary 分页获取Sport_score列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Sport_scoreSearch true "分页获取Sport_score列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sport_score/getSport_scoreList [get]
func GetSport_scoreList(c *gin.Context) {
	var pageInfo request.Sport_scoreSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSport_scoreInfoList(pageInfo); err != nil {
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
