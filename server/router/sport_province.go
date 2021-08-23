package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitSportProvinceRouter 初始化 SportProvince 路由信息
func InitSportProvinceRouter(Router *gin.RouterGroup) {
	SportProvinceRouter := Router.Group("sportProvince").Use(middleware.OperationRecord())
	{
		SportProvinceRouter.POST("createSportProvince", v1.CreateSportProvince)   // 新建SportProvince
		SportProvinceRouter.DELETE("deleteSportProvince", v1.DeleteSportProvince) // 删除SportProvince
		SportProvinceRouter.DELETE("deleteSportProvinceByIds", v1.DeleteSportProvinceByIds) // 批量删除SportProvince
		SportProvinceRouter.PUT("updateSportProvince", v1.UpdateSportProvince)    // 更新SportProvince
		SportProvinceRouter.GET("findSportProvince", v1.FindSportProvince)        // 根据ID获取SportProvince
		SportProvinceRouter.GET("getSportProvinceList", v1.GetSportProvinceList)  // 获取SportProvince列表
	}
}
