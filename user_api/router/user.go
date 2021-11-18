package router

import (
	"github.com/gin-gonic/gin"
	"mxshop_api/user_api/api"
	"mxshop_api/user_api/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
	}
}
