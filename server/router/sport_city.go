package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitSportCityRouter 初始化 SportCity 路由信息
func InitSportCityRouter(Router *gin.RouterGroup) {
	SportCityRouter := Router.Group("sportCity").Use(middleware.OperationRecord())
	{
		SportCityRouter.POST("createSportCity", v1.CreateSportCity)   // 新建SportCity
		SportCityRouter.DELETE("deleteSportCity", v1.DeleteSportCity) // 删除SportCity
		SportCityRouter.DELETE("deleteSportCityByIds", v1.DeleteSportCityByIds) // 批量删除SportCity
		SportCityRouter.PUT("updateSportCity", v1.UpdateSportCity)    // 更新SportCity
		SportCityRouter.GET("findSportCity", v1.FindSportCity)        // 根据ID获取SportCity
		SportCityRouter.GET("getSportCityList", v1.GetSportCityList)  // 获取SportCity列表
	}
}
