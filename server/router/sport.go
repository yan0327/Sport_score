package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitSportRouter 初始化 Sport 路由信息
func InitSportRouter(Router *gin.RouterGroup) {
	SportRouter := Router.Group("sport").Use(middleware.OperationRecord())
	{
		SportRouter.POST("createSport", v1.CreateSport)   // 新建Sport
		SportRouter.DELETE("deleteSport", v1.DeleteSport) // 删除Sport
		SportRouter.DELETE("deleteSportByIds", v1.DeleteSportByIds) // 批量删除Sport
		SportRouter.PUT("updateSport", v1.UpdateSport)    // 更新Sport
		SportRouter.GET("findSport", v1.FindSport)        // 根据ID获取Sport
		SportRouter.GET("getSportList", v1.GetSportList)  // 获取Sport列表
	}
}
