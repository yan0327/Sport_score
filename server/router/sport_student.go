package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitSportStudentRouter 初始化 SportStudent 路由信息
func InitSportStudentRouter(Router *gin.RouterGroup) {
	SportStudentRouter := Router.Group("sportStudent").Use(middleware.OperationRecord())
	{
		SportStudentRouter.POST("createSportStudent", v1.CreateSportStudent)   // 新建SportStudent
		SportStudentRouter.DELETE("deleteSportStudent", v1.DeleteSportStudent) // 删除SportStudent
		SportStudentRouter.DELETE("deleteSportStudentByIds", v1.DeleteSportStudentByIds) // 批量删除SportStudent
		SportStudentRouter.PUT("updateSportStudent", v1.UpdateSportStudent)    // 更新SportStudent
		SportStudentRouter.GET("findSportStudent", v1.FindSportStudent)        // 根据ID获取SportStudent
		SportStudentRouter.GET("getSportStudentList", v1.GetSportStudentList)  // 获取SportStudent列表
	}
}
