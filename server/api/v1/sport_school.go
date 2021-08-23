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

// CreateSportSchool 创建SportSchool
// @Tags SportSchool
// @Summary 创建SportSchool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportSchool true "创建SportSchool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportSchool/createSportSchool [post]
func CreateSportSchool(c *gin.Context) {
	var sportSchool model.SportSchool
	_ = c.ShouldBindJSON(&sportSchool)
	if err := service.CreateSportSchool(sportSchool); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSportSchool 删除SportSchool
// @Tags SportSchool
// @Summary 删除SportSchool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportSchool true "删除SportSchool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportSchool/deleteSportSchool [delete]
func DeleteSportSchool(c *gin.Context) {
	var sportSchool model.SportSchool
	_ = c.ShouldBindJSON(&sportSchool)
	if err := service.DeleteSportSchool(sportSchool); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSportSchoolByIds 批量删除SportSchool
// @Tags SportSchool
// @Summary 批量删除SportSchool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SportSchool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sportSchool/deleteSportSchoolByIds [delete]
func DeleteSportSchoolByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSportSchoolByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSportSchool 更新SportSchool
// @Tags SportSchool
// @Summary 更新SportSchool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportSchool true "更新SportSchool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sportSchool/updateSportSchool [put]
func UpdateSportSchool(c *gin.Context) {
	var sportSchool model.SportSchool
	_ = c.ShouldBindJSON(&sportSchool)
	if err := service.UpdateSportSchool(sportSchool); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSportSchool 用id查询SportSchool
// @Tags SportSchool
// @Summary 用id查询SportSchool
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportSchool true "用id查询SportSchool"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sportSchool/findSportSchool [get]
func FindSportSchool(c *gin.Context) {
	var sportSchool model.SportSchool
	_ = c.ShouldBindQuery(&sportSchool)
	if err, resportSchool := service.GetSportSchool(sportSchool.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resportSchool": resportSchool}, c)
	}
}

// GetSportSchoolList 分页获取SportSchool列表
// @Tags SportSchool
// @Summary 分页获取SportSchool列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SportSchoolSearch true "分页获取SportSchool列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportSchool/getSportSchoolList [get]
func GetSportSchoolList(c *gin.Context) {
	var pageInfo request.SportSchoolSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSportSchoolInfoList(pageInfo); err != nil {
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
