package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitSportSchoolRouter 初始化 SportSchool 路由信息
func InitSportSchoolRouter(Router *gin.RouterGroup) {
	SportSchoolRouter := Router.Group("sportSchool").Use(middleware.OperationRecord())
	{
		SportSchoolRouter.POST("createSportSchool", v1.CreateSportSchool)   // 新建SportSchool
		SportSchoolRouter.DELETE("deleteSportSchool", v1.DeleteSportSchool) // 删除SportSchool
		SportSchoolRouter.DELETE("deleteSportSchoolByIds", v1.DeleteSportSchoolByIds) // 批量删除SportSchool
		SportSchoolRouter.PUT("updateSportSchool", v1.UpdateSportSchool)    // 更新SportSchool
		SportSchoolRouter.GET("findSportSchool", v1.FindSportSchool)        // 根据ID获取SportSchool
		SportSchoolRouter.GET("getSportSchoolList", v1.GetSportSchoolList)  // 获取SportSchool列表
	}
}
