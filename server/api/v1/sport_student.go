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

// CreateSportStudent 创建SportStudent
// @Tags SportStudent
// @Summary 创建SportStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudent true "创建SportStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportStudent/createSportStudent [post]
func CreateSportStudent(c *gin.Context) {
	var sportStudent model.SportStudent
	_ = c.ShouldBindJSON(&sportStudent)
	if err := service.CreateSportStudent(sportStudent); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSportStudent 删除SportStudent
// @Tags SportStudent
// @Summary 删除SportStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudent true "删除SportStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sportStudent/deleteSportStudent [delete]
func DeleteSportStudent(c *gin.Context) {
	var sportStudent model.SportStudent
	_ = c.ShouldBindJSON(&sportStudent)
	if err := service.DeleteSportStudent(sportStudent); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSportStudentByIds 批量删除SportStudent
// @Tags SportStudent
// @Summary 批量删除SportStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SportStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sportStudent/deleteSportStudentByIds [delete]
func DeleteSportStudentByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSportStudentByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSportStudent 更新SportStudent
// @Tags SportStudent
// @Summary 更新SportStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudent true "更新SportStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sportStudent/updateSportStudent [put]
func UpdateSportStudent(c *gin.Context) {
	var sportStudent model.SportStudent
	_ = c.ShouldBindJSON(&sportStudent)
	if err := service.UpdateSportStudent(sportStudent); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSportStudent 用id查询SportStudent
// @Tags SportStudent
// @Summary 用id查询SportStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SportStudent true "用id查询SportStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sportStudent/findSportStudent [get]
func FindSportStudent(c *gin.Context) {
	var sportStudent model.SportStudent
	_ = c.ShouldBindQuery(&sportStudent)
	if err, resportStudent := service.GetSportStudent(sportStudent.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resportStudent": resportStudent}, c)
	}
}

// GetSportStudentList 分页获取SportStudent列表
// @Tags SportStudent
// @Summary 分页获取SportStudent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SportStudentSearch true "分页获取SportStudent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sportStudent/getSportStudentList [get]
func GetSportStudentList(c *gin.Context) {
	var pageInfo request.SportStudentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSportStudentInfoList(pageInfo); err != nil {
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
