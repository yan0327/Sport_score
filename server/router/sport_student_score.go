package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitSportStudentScoreRouter 初始化 SportStudentScore 路由信息
func InitSportStudentScoreRouter(Router *gin.RouterGroup) {
	SportStudentScoreRouter := Router.Group("sportStudentScore").Use(middleware.OperationRecord())
	{
		SportStudentScoreRouter.POST("createSportStudentScore", v1.CreateSportStudentScore)   // 新建SportStudentScore
		SportStudentScoreRouter.DELETE("deleteSportStudentScore", v1.DeleteSportStudentScore) // 删除SportStudentScore
		SportStudentScoreRouter.DELETE("deleteSportStudentScoreByIds", v1.DeleteSportStudentScoreByIds) // 批量删除SportStudentScore
		SportStudentScoreRouter.PUT("updateSportStudentScore", v1.UpdateSportStudentScore)    // 更新SportStudentScore
		SportStudentScoreRouter.GET("findSportStudentScore", v1.FindSportStudentScore)        // 根据ID获取SportStudentScore
		SportStudentScoreRouter.GET("getSportStudentScoreList", v1.GetSportStudentScoreList)  // 获取SportStudentScore列表
	}
}
