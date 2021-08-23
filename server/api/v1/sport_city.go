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

// CreateSportCity 创建SportCity
// @Tags SportCity
// @Summary 创建SportCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportCity true "创建SportCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportCity/createSportCity [post]
func CreateSportCity(c *gin.Context) {
	var sportCity model.SportCity
	_ = c.ShouldBindJSON(&sportCity)
	if err := service.CreateSportCity(sportCity); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSportCity 删除SportCity
// @Tags SportCity
// @Summary 删除SportCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportCity true "删除SportCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportCity/deleteSportCity [delete]
func DeleteSportCity(c *gin.Context) {
	var sportCity model.SportCity
	_ = c.ShouldBindJSON(&sportCity)
	if err := service.DeleteSportCity(sportCity); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSportCityByIds 批量删除SportCity
// @Tags SportCity
// @Summary 批量删除SportCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SportCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sportCity/deleteSportCityByIds [delete]
func DeleteSportCityByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSportCityByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSportCity 更新SportCity
// @Tags SportCity
// @Summary 更新SportCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportCity true "更新SportCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sportCity/updateSportCity [put]
func UpdateSportCity(c *gin.Context) {
	var sportCity model.SportCity
	_ = c.ShouldBindJSON(&sportCity)
	if err := service.UpdateSportCity(sportCity); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSportCity 用id查询SportCity
// @Tags SportCity
// @Summary 用id查询SportCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportCity true "用id查询SportCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sportCity/findSportCity [get]
func FindSportCity(c *gin.Context) {
	var sportCity model.SportCity
	_ = c.ShouldBindQuery(&sportCity)
	if err, resportCity := service.GetSportCity(sportCity.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resportCity": resportCity}, c)
	}
}

// GetSportCityList 分页获取SportCity列表
// @Tags SportCity
// @Summary 分页获取SportCity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SportCitySearch true "分页获取SportCity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportCity/getSportCityList [get]
func GetSportCityList(c *gin.Context) {
	var pageInfo request.SportCitySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSportCityInfoList(pageInfo); err != nil {
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
