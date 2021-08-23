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

// CreateSportProvince 创建SportProvince
// @Tags SportProvince
// @Summary 创建SportProvince
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportProvince true "创建SportProvince"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportProvince/createSportProvince [post]
func CreateSportProvince(c *gin.Context) {
	var sportProvince model.SportProvince
	_ = c.ShouldBindJSON(&sportProvince)
	if err := service.CreateSportProvince(sportProvince); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSportProvince 删除SportProvince
// @Tags SportProvince
// @Summary 删除SportProvince
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportProvince true "删除SportProvince"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportProvince/deleteSportProvince [delete]
func DeleteSportProvince(c *gin.Context) {
	var sportProvince model.SportProvince
	_ = c.ShouldBindJSON(&sportProvince)
	if err := service.DeleteSportProvince(sportProvince); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSportProvinceByIds 批量删除SportProvince
// @Tags SportProvince
// @Summary 批量删除SportProvince
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SportProvince"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sportProvince/deleteSportProvinceByIds [delete]
func DeleteSportProvinceByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSportProvinceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSportProvince 更新SportProvince
// @Tags SportProvince
// @Summary 更新SportProvince
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportProvince true "更新SportProvince"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sportProvince/updateSportProvince [put]
func UpdateSportProvince(c *gin.Context) {
	var sportProvince model.SportProvince
	_ = c.ShouldBindJSON(&sportProvince)
	if err := service.UpdateSportProvince(sportProvince); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSportProvince 用id查询SportProvince
// @Tags SportProvince
// @Summary 用id查询SportProvince
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportProvince true "用id查询SportProvince"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sportProvince/findSportProvince [get]
func FindSportProvince(c *gin.Context) {
	var sportProvince model.SportProvince
	_ = c.ShouldBindQuery(&sportProvince)
	if err, resportProvince := service.GetSportProvince(sportProvince.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resportProvince": resportProvince}, c)
	}
}

// GetSportProvinceList 分页获取SportProvince列表
// @Tags SportProvince
// @Summary 分页获取SportProvince列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SportProvinceSearch true "分页获取SportProvince列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportProvince/getSportProvinceList [get]
func GetSportProvinceList(c *gin.Context) {
	var pageInfo request.SportProvinceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSportProvinceInfoList(pageInfo); err != nil {
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
