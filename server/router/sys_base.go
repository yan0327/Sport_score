package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("login2", v1.Login2)
		BaseRouter.POST("captcha", v1.Captcha)
		BaseRouter.POST("information", v1.Information)
		BaseRouter.POST("createInformation", v1.CreateInformation)
	}
	return BaseRouter
}
