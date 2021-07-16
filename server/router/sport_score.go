package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitSport_scoreRouter 初始化 Sport_score 路由信息
func InitSport_scoreRouter(Router *gin.RouterGroup) {
	Sport_scoreRouter := Router.Group("sport_score").Use(middleware.OperationRecord())
	{
		Sport_scoreRouter.POST("createSport_score", v1.CreateSport_score)   // 新建Sport_score
		Sport_scoreRouter.DELETE("deleteSport_score", v1.DeleteSport_score) // 删除Sport_score
		Sport_scoreRouter.DELETE("deleteSport_scoreByIds", v1.DeleteSport_scoreByIds) // 批量删除Sport_score
		Sport_scoreRouter.PUT("updateSport_score", v1.UpdateSport_score)    // 更新Sport_score
		Sport_scoreRouter.GET("findSport_score", v1.FindSport_score)        // 根据ID获取Sport_score
		Sport_scoreRouter.GET("getSport_scoreList", v1.GetSport_scoreList)  // 获取Sport_score列表
	}
}
